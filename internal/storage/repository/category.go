package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/google/uuid"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (c *CategoryRepo) CreateCategory(req *pb.CreateCategoryReq) (*pb.Void, error) {

	id := uuid.NewString()
	query := `INSERT INTO category 
	(id, name, description, created_at) VALUES ($1, $2, $3, $4)`

	_, err := c.db.Exec(query, id, req.Name, req.Description, time.Now())
	if err != nil {
		log.Println("Error while creating category", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (c *CategoryRepo) GetCategory(req *pb.GetById) (*pb.Categories, error) {

	query := `SELECT
		id,
		name,
		description,
		created_at
	FROM
		category
	WHERE
		id = $1
	AND
		deleted_at = 0
	`
	row := c.db.QueryRow(query, req.Id)
	var category pb.Categories
	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.CreatedAt)

	if err != nil {
		log.Println("Error while getting category", err)
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepo) ListCategories(req *pb.ListAllCategoriesReq) (*pb.ListCategoriesRes, error) {

	query := `SELECT
		id,
		name,
		description,
		created_at
	FROM
		category
	WHERE
		deleted_at = 0
	`
	var args []interface{}
	argCount := 1
	filters := []string{}
	if req.Name != "" {
		filters = append(filters, fmt.Sprintf("name = $%d", argCount))
		args = append(args, req.Name)
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
		return nil, err
	}
	defer rows.Close()
	var categories []*pb.Categories
	for rows.Next() {
		var category pb.Categories
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.CreatedAt)
		if err != nil {
			log.Println("Error while getting categories", err)
			return nil, err
		}
		categories = append(categories, &category)
	}
	count := len(categories)
	return &pb.ListCategoriesRes{Categories: categories, Count: int32(count)}, nil
}

func (c *CategoryRepo) UpdateCategory(req *pb.UpdateCategoryReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.Name != "" && req.Body.Name != "string" {
		args = append(args, req.Body.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}
	if req.Body.Description != "" && req.Body.Description != "string" {
		args = append(args, req.Body.Description)
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE category SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := c.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating categories", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (c *CategoryRepo) DeleteCategory(req *pb.GetById) (*pb.Void, error) {
	query := `
	UPDATE 
		category
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	AND
		deleted_at = 0
	`
	_, err := c.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting category", err)
		return nil, err
	}

	return &pb.Void{}, nil

}
