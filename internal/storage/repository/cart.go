package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/google/uuid"
)

type CartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (c *CartRepo) CreateCart(req *pb.CreateCartReq) (*pb.Void, error) {
	id := uuid.NewString()
	query :=
		`INSERT INTO carts
		(id, cart_number, cart_name, amount,user_id) 
		VALUES ($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(query, id, req.CartName, req.Amount, req.CartNumber, req.UserId)
	if err != nil {
		log.Println("Error while creating Cart", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (c *CartRepo) GetCart(req *pb.GetById) (*pb.Cart, error) {
	var cart pb.Cart
	query := `SELECT
		c.id
		c.cart_number,
		c.cart_name,
		c.amount,
		c.created_at,
		u.id,
		u.username,
		u.email,
		u.fullname,
		u.date_of_birth
	FROM
		carts c
	LEFT JOIN
		users u
	ON
		c.user_id = u.id
	WHERE
		c.id = $1
	AND
		c.deleted_at = 0`

	row := c.db.QueryRow(query, req.Id)
	err := row.Scan(
		&cart.Id,
		&cart.CartNumber,
		&cart.Amount,
		&cart.CartName,
		&cart.CreatedAt,
		&cart.User.Id,
		&cart.User.FullName,
		&cart.User.Email,
		&cart.User.Username,
		&cart.User.DateOfBirth,
	)
	if err != nil {
		log.Println("Error while getting Cart", err)
		return nil, err
	}
	return &cart, nil
}

func (c *CartRepo) ListAllCarts(req *pb.ListAllCartReq) (*pb.ListAllCartRes, error) {

	query := `SELECT
		c.id
		c.cart_number,
		c.cart_name,
		c.amount,
		c.created_at,
		u.id,
		u.username,
		u.email,
		u.fullname,
		u.date_of_birth
	FROM
		carts c
	LEFT JOIN
		users u
	ON
		c.user_id = u.id
	WHERE
		c.deleted_at = 0`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.CartName != "" {
		filters = append(filters, fmt.Sprintf("cart_name = $%d", argCount))
		args = append(args, req.CartName)
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

	rows, err := c.db.Query(query, args...)
	if err != nil {
		log.Println("Error while listing Carts", err)
		return nil, err
	}
	defer rows.Close()
	carts := []*pb.Cart{}
	for rows.Next() {
		var cart pb.Cart
		err := rows.Scan(
			&cart.Id,
			&cart.CartName,
			&cart.Amount,
			&cart.CartNumber,
			&cart.CreatedAt,
			&cart.User.Id,
			&cart.User.Id,
			&cart.User.FullName,
			&cart.User.Email,
			&cart.User.Username,
			&cart.User.DateOfBirth,
		)
		if err != nil {
			log.Println("Error while scanning Cart", err)
			return nil, err
		}
		carts = append(carts, &cart)
	}
	count := len(carts)

	return &pb.ListAllCartRes{
		Carts:  carts,
		Count: int32(count),
	}, nil
}
