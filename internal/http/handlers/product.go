package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Product
// @Description Create a new Product
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Product body pb.CreateProductReq true "Product data"
// @Success 200 {string} string "message":"Product created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/create [post]
func (h *Handlers) CreateProduct(c *gin.Context) {
	var req pb.CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("create-product", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Product created successfully"})
}

// @Summary Get Product
// @Description Get an Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} pb.Products "Product data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/{id} [get]
func (h *Handlers) GetProduct(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Product.GetProduct(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Product
// @Description Update an existing Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Product body pb.UpdateProductReq true "Product update data"
// @Success 200 {string} string "message":"Product updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/update/{id} [put]
func (h *Handlers) UpdateProduct(c *gin.Context) {
	var req pb.UpdateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("update-product", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

// @Summary List Products
// @Description List Products with filters
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query int false "Name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListAllProductsRes "List of Products"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/list [get]
func (h *Handlers) ListProducts(c *gin.Context) {
	var filter pb.ListAllProductsReq

	name := c.Query("name")
	filter.Name = name

	f := pb.Pagination{}
	filter.Pagination = &f

	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Pagination.Limit = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Pagination.Offset = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.Product.ListAllProducts(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Product
// @Description Delete an Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {string} string "message":"Product deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/delete/{id} [delete]
func (h *Handlers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Product.DeleteProduct(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

// @Summary Get Product By Product ID
// @Description Get an Product by Product ID
// @Tags Product
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param        min_price path int true "min_price"
// @Param        max_price path int true "max_price"
// @Success 200 {object} pb.ListAllProductsRes "Product data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/by-range [get]
func (h *Handlers) GetProductsByPriceRange(c *gin.Context) {
	req := pb.GetProductsByPriceRangeReq{}
	minpriceSTR := c.Query("min_price")
	if minpriceSTR == "" {
		minpriceSTR = "0"
	}
	minprice, err := strconv.Atoi(minpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid min_price parameter: " + err.Error(),
		})
		return
	}

	maxpriceSTR := c.Query("max_price")
	if maxpriceSTR == "" {
		maxpriceSTR = "0"
	}
	maxprice, err := strconv.Atoi(maxpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid max_price parameter: " + err.Error(),
		})
		return
	}

	req.MaxPrice = float32(maxprice)
	req.MinPrice = float32(minprice)

	res, err := h.Product.GetProductsByPriceRange(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)

}
