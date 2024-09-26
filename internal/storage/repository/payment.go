package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/google/uuid"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{
		db: db,
	}
}

func (pr *PaymentRepo) CreatePayment(req *pb.CreatePaymentReq) (*pb.Void, error) {
	id := uuid.NewString()
	tx, err := pr.db.Begin() 
	if err != nil {
		log.Println("Error starting transaction", err)
		return nil, err
	}

	var currentAmount int32
	cartQuery := `SELECT amount FROM cart WHERE card_id = $1`
	err = tx.QueryRow(cartQuery, req.CardId).Scan(&currentAmount)
	if err != nil {
		log.Println("Error fetching current amount from cart", err)
		tx.Rollback() 
		return nil, err
	}

	if currentAmount < req.Amount {
		log.Println("Insufficient amount in the cart")
		tx.Rollback()
		return nil, fmt.Errorf("insufficient amount in the cart")
	}
	newAmount := currentAmount - req.Amount

	updateCartQuery := `UPDATE cart SET amount = $1 WHERE card_id = $2`
	_, err = tx.Exec(updateCartQuery, newAmount, req.CardId)
	if err != nil {
		log.Println("Error updating cart amount", err)
		tx.Rollback() 
		return nil, err
	}

	insertPaymentQuery :=
		`INSERT INTO payments
		(id, payment_method, amount, order_id, card_id, status) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(insertPaymentQuery, id, req.PaymentMethod, req.Amount, req.OrderID, req.CardId, req.Status)
	if err != nil {
		log.Println("Error inserting payment", err)
		tx.Rollback() 
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction", err)
		return nil, err
	}

	return &pb.Void{}, nil
}


func (pr *PaymentRepo) GetPayment(req *pb.GetById) (*pb.Payment, error) {
	var payment pb.Payment
	query := `SELECT
		p.id,
		p.payment_method, 
		p.amount, 
		p.status,
		p.cart_id,
		p.created_at,
		o.id,
		o.user_id,
		o.product_id,
		o.total_price, 
		o.quantity, 
		o.status, 
		o.created_at
	FROM
		payments
	LEFT JOIN
		orders o
	ON
		p.order_id = o.id
	WHERE
		p.id = $1
	AND
		p.deleted_at = 0`

	row := pr.db.QueryRow(query, req.Id)
	err := row.Scan(
		&payment.Id,
		&payment.PaymentMethod,
		&payment.Amount,
		&payment.Status,
		&payment.CartId,
		&payment.CreatedAt,
		&payment.Order.Id,
		&payment.Order.UserID,
		&payment.Order.ProductID,
		&payment.Order.TotalPrice,
		&payment.Order.Quantity,
		&payment.Order.Status,
		&payment.Order.CreatedAt,
	)
	if err != nil {
		log.Println("Error while getting payment", err)
		return nil, err
	}
	return &payment, nil
}

func (pr *PaymentRepo) ListPayments(req *pb.ListPaymentsReq) (*pb.ListPaymentsRes, error) {

	query := `SELECT
		p.id,
		p.payment_method, 
		p.amount, 
		p.status,
		p.cart_id,
		p.created_at,
		o.id,
		o.user_id,
		o.product_id,
		o.total_price, 
		o.quantity, 
		o.status, 
		o.created_at
	FROM
		payments
	LEFT JOIN
		orders o
	ON
		p.order_id = o.id
	WHERE
		p.deleted_at = 0`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.Status != "" {
		filters = append(filters, fmt.Sprintf("status = $%d", argCount))
		args = append(args, req.Status)
		argCount++
	}

	if req.PaymentMethod != "" {
		filters = append(filters, fmt.Sprintf("payment_method = $%d", argCount))
		args = append(args, req.PaymentMethod)
		argCount++
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}
	if req.Pagination != nil {
		if req.Pagination.Limit > 0 {
			query += fmt.Sprintf(" LIMIT $%d", argCount)
			args = append(args, req.Pagination.Limit)
			argCount++
		}
		if req.Pagination.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Pagination.Offset)
			argCount++
		}
	}

	rows, err := pr.db.Query(query, args...)
	if err != nil {
		log.Println("Error while listing payments", err)
		return nil, err
	}
	defer rows.Close()
	payments := []*pb.Payment{}
	for rows.Next() {
		var payment pb.Payment
		err := rows.Scan(
			&payment.Id,
			&payment.PaymentMethod,
			&payment.Amount,
			&payment.Status,
			&payment.CartId,
			&payment.CreatedAt,
			&payment.Order.Id,
			&payment.Order.UserID,
			&payment.Order.ProductID,
			&payment.Order.TotalPrice,
			&payment.Order.Quantity,
			&payment.Order.Status,
			&payment.Order.CreatedAt,
		)
		if err != nil {
			log.Println("Error while scanning payment", err)
			return nil, err
		}
		payments = append(payments, &payment)
	}
	count := len(payments)

	return &pb.ListPaymentsRes{
		Payment: payments,
		Count:   int32(count),
	}, nil
}
