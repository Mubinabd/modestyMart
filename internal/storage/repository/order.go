package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/google/uuid"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (o *OrderRepo) CreateOrder(req *pb.CreateOrderReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO orders (id,user_id, product_id, quantity, total_price, status)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	_, err := o.db.Exec(query, id, req.UserID, req.ProductID, req.Quantity, req.TotalPrice, req.Status)
	if err != nil {
		log.Println("Error while creating order", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (o *OrderRepo) UpdateOrder(req *pb.UpdateOrderReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Order.UserID != "" && req.Order.UserID != "string" {
		args = append(args, req.Order.UserID)
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)))
	}
	if req.Order.ProductID != "" && req.Order.ProductID != "string" {
		args = append(args, req.Order.ProductID)
		conditions = append(conditions, fmt.Sprintf("product_id = $%d", len(args)))
	}
	if req.Order.TotalPrice != 0 {
		args = append(args, req.Order.TotalPrice)
		conditions = append(conditions, fmt.Sprintf("total_price = $%d", len(args)))
	}
	if req.Order.Status != "" && req.Order.Status != "string" {
		args = append(args, req.Order.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}
	if req.Order.Quantity != 0 {
		args = append(args, req.Order.Quantity)
		conditions = append(conditions, fmt.Sprintf("quantity = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE orders SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := o.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating order", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (o *OrderRepo) DeleteOrder(req *pb.GetById) (*pb.Void, error) {
	query := `
	UPDATE 
		orders
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	AND
		deleted_at = 0
	`
	_, err := o.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting order", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (o *OrderRepo) GetOrder(req *pb.GetById) (*pb.Orders, error) {
	query := `
	SELECT
		o.id, 
		o.total_price, 
		o.status, 
		o.quantity, 
		o.created_at,
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.date_of_birth,
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.category_id,
		p.image_url,
		p.created_at
	FROM
		orders o
	LEFT JOIN
		users u
	ON
		o.user_id = u.id
	LEFT JOIN
		products p
	ON
		o.product_id = p.id
	WHERE
		o.id = $1
	AND
		o.deleted_at = 0
	`
	var order pb.Orders
	order.User = &pb.UserRes{}
	order.Product = &pb.ProductsRes{}
	err := o.db.QueryRow(query, req.Id).
		Scan(
			&order.Id,
			&order.TotalPrice,
			&order.Status,
			&order.Quantity,
			&order.CreatedAt,
			&order.User.Id,
			&order.User.Username,
			&order.User.Email,
			&order.User.FullName,
			&order.User.DateOfBirth,
			&order.Product.Id,
			&order.Product.Name,
			&order.Product.Description,
			&order.Product.Price,
			&order.Product.Stock,
			&order.Product.CategoryID,
			&order.Product.ImageUrl,
			&order.Product.CreatedAt)
	if err != nil {
		log.Println("Error while getting order", err)
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) ListAllOrders(req *pb.ListOrdersReq) (*pb.ListOrdersRes, error) {

	query := `
	SELECT
		o.id, 
		o.total_price, 
		o.status, 
		o.quantity, 
		o.created_at,
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.date_of_birth,
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.category_id,
		p.image_url,
		p.created_at
	FROM
		orders o
	LEFT JOIN
		users u
	ON
		o.user_id = u.id
	LEFT JOIN
		products p
	ON
		o.product_id = p.id
	WHERE
		o.deleted_at = 0
	`
	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.Status != "" {
		filters = append(filters, fmt.Sprintf("status = $%d", argCount))
		args = append(args, req.Status)
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
	rows, err := o.db.Query(query, args...)
	if err != nil {
		log.Println("Error while listing orders", err)
		return nil, err
	}
	defer rows.Close()
	var orders []*pb.Orders
	for rows.Next() {
		var order pb.Orders
		order.User = &pb.UserRes{}
		order.Product = &pb.ProductsRes{}
		err := rows.Scan(
			&order.Id,
			&order.TotalPrice,
			&order.Status,
			&order.Quantity,
			&order.CreatedAt,
			&order.User.Id,
			&order.User.Username,
			&order.User.Email,
			&order.User.FullName,
			&order.User.DateOfBirth,
			&order.Product.Id,
			&order.Product.Name,
			&order.Product.Description,
			&order.Product.Price,
			&order.Product.Stock,
			&order.Product.CategoryID,
			&order.Product.ImageUrl,
			&order.Product.CreatedAt)
		if err != nil {
			log.Println("Error while scanning orders", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	count := len(orders)
	return &pb.ListOrdersRes{
		Orders: orders,
		Count:  int32(count),
	}, nil
}

func (o *OrderRepo) GetOrderByProductID(req *pb.OrderByProductId) (*pb.GetOrdersRes, error) {
	// Ensure ProductID is not empty
	if req.ProductID == "" {
		log.Println("Product ID cannot be empty")
		return nil, errors.New("invalid product ID")
	}

	query := `
	SELECT
		o.id, 
		o.total_price, 
		o.status, 
		o.quantity, 
		o.created_at,
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.date_of_birth,
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.category_id,
		p.image_url,
		p.created_at
	FROM		
		orders o
	LEFT JOIN
		users u
	ON
		o.user_id = u.id
	LEFT JOIN
		products p
	ON
		o.product_id = p.id
	WHERE
		p.id = $1
	AND
		o.deleted_at IS NULL
	`

	rows, err := o.db.Query(query, req.ProductID)
	if err != nil {
		log.Println("Error while listing orders:", err)
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Orders
	for rows.Next() {
		var order pb.Orders
		order.User = &pb.UserRes{}
		order.Product = &pb.ProductsRes{}
		err := rows.Scan(
			&order.Id,
			&order.TotalPrice,
			&order.Status,
			&order.Quantity,
			&order.CreatedAt,
			&order.User.Id,
			&order.User.Username,
			&order.User.Email,
			&order.User.FullName,
			&order.User.DateOfBirth,
			&order.Product.Id,
			&order.Product.Name,
			&order.Product.Description,
			&order.Product.Price,
			&order.Product.Stock,
			&order.Product.CategoryID,
			&order.Product.ImageUrl,
			&order.Product.CreatedAt)
		if err != nil {
			log.Println("Error while scanning orders:", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	count := len(orders)
	return &pb.GetOrdersRes{
		Orders: orders,
		Count:  int32(count),
	}, nil
}
