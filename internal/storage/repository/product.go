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

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) CreateProduct(req *pb.CreateProductReq) (*pb.Void, error) {
	id := uuid.NewString()

	query := `INSERT INTO products 
		(id, name, description, price, stock, category_id,image_url) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := p.db.Exec(query, id, req.Name, req.Description, req.Price, req.Stock, req.CategoryID, req.ImageUrl)
	if err != nil {
		log.Println("Error while creating product", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (p *ProductRepo) UpdateProduct(req *pb.UpdateProductReq) (*pb.Void, error) {
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
	if req.Body.CategoryID != "" && req.Body.CategoryID != "string" {
		args = append(args, req.Body.CategoryID)
		conditions = append(conditions, fmt.Sprintf("category_id = $%d", len(args)))
	}
	if req.Body.ImageUrl != "" && req.Body.ImageUrl != "string" {
		args = append(args, req.Body.ImageUrl)
		conditions = append(conditions, fmt.Sprintf("image_url = $%d", len(args)))
	}
	if req.Body.Price != 0 {
		args = append(args, req.Body.Price)
		conditions = append(conditions, fmt.Sprintf("price = $%d", len(args)))
	}
	if req.Body.Stock != 0 {
		args = append(args, req.Body.Stock)
		conditions = append(conditions, fmt.Sprintf("stock = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE products SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := p.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating products", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (p *ProductRepo) DeleteProduct(req *pb.GetById) (*pb.Void, error) {
	query := `
		UPDATE
			products
		SET
			deleted_at = extract(epoch from now())
		WHERE
			id = $1
		`

	_, err := p.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting product", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *ProductRepo) GetProduct(req *pb.GetById) (*pb.Products, error) {

	query := `
	SELECT
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.image_url,
		p.created_at,
		c.id,
		c.name,
		c.description,
		c.created_at
	FROM
		products p
	LEFT JOIN
		category c
	ON
		p.category_id = c.id
	WHERE
		p.id = $1
	AND
		p.deleted_at = 0
	`
	row := p.db.QueryRow(query, req.Id)
	var product pb.Products
	product.Category = &pb.Categories{}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.ImageUrl,
		&product.CreatedAt,
		&product.Category.Id,
		&product.Category.Name,
		&product.Category.Description,
		&product.Category.CreatedAt,
	)
	if err != nil {
		log.Println("Error while fetching product", err)
		return nil, err
	}
	return &product, nil
}
func (p *ProductRepo) ListAllProducts(req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error) {

	query := `
	SELECT
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.image_url,
		p.created_at,
		c.id,
		c.name,
		c.description,
		c.created_at
	FROM
		products p
	LEFT JOIN
		category c
	ON
		p.category_id = c.id
	WHERE
		p.deleted_at = 0
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

	rows, err := p.db.Query(query, args...)
	if err != nil {
		log.Println("Error while fetching products", err)
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Products
	for rows.Next() {
		var product pb.Products
		product.Category = &pb.Categories{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.ImageUrl,
			&product.CreatedAt,
			&product.Category.Id,
			&product.Category.Name,
			&product.Category.Description,
			&product.Category.CreatedAt,
		)
		if err != nil {
			log.Println("Error while fetching products", err)
			return nil, err
		}
		products = append(products, &product)
	}
	count := len(products)
	return &pb.ListAllProductsRes{Products: products, Count: int32(count)}, nil
}

func (p *ProductRepo) GetCategory(req *pb.GetCategoryReq) (*pb.GetCategoryRes, error) {
	query := `
	SELECT
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.image_url,
		p.created_at,
		c.id,
		c.name,
		c.description,
		c.created_at
	FROM
		products p
	LEFT JOIN
		category c
	ON
		p.category_id = c.id
	WHERE
		c.id = $1
	AND
		p.deleted_at = 0
	`

	row := p.db.QueryRow(query, req.CategoryID)

    var product pb.GetCategoryRes
    product.Products = &pb.Products{}
    product.Products.Category = &pb.Categories{}

	// Perform the scan
	err := row.Scan(
		&product.Products.Id,
		&product.Products.Name,
		&product.Products.Description,
		&product.Products.Price,
		&product.Products.Stock,
		&product.Products.ImageUrl,
		&product.Products.CreatedAt,
		&product.Products.Category.Id,
		&product.Products.Category.Name,
		&product.Products.Category.Description,
		&product.Products.Category.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Category not found", err)
			return nil, sql.ErrNoRows // Let the handler handle this as a 404
		}
		log.Println("Error while fetching product", err)
		return nil, err
	}
	return &product, nil
}


func (p *ProductRepo) GetProductsByPriceRange(req *pb.GetProductsByPriceRangeReq) (*pb.ListAllProductsRes, error) {
	query := `
	SELECT
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.image_url,
		p.created_at,
		c.id,
		c.name,
		c.description,
		c.created_at
	FROM
		products p
	LEFT JOIN
		category c
	ON
		p.category_id = c.id
	WHERE
		p.deleted_at = 0
	`
	var args []interface{}
	filters := []string{}

	if req.MinPrice != 0 {
		args = append(args, req.MinPrice)
		filters = append(filters, fmt.Sprintf("p.price >= $%d", len(args)))
	}

	if req.MaxPrice != 0 {
		args = append(args, req.MaxPrice)
		filters = append(filters, fmt.Sprintf("p.price <= $%d", len(args)))
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		log.Println("Error while fetching products", err)
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Products
	for rows.Next() {
		var product pb.Products
		product.Category = &pb.Categories{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.ImageUrl,
			&product.CreatedAt,
			&product.Category.Id,
			&product.Category.Name,
			&product.Category.Description,
			&product.Category.CreatedAt,
		)
		if err != nil {
			log.Println("Error while scanning product data", err)
			return nil, err
		}
		products = append(products, &product)
	}

	count := len(products)
	return &pb.ListAllProductsRes{Products: products, Count: int32(count)}, nil
}
