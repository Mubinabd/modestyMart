package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Category
// @Description Create a new Category
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Category body pb.CreateCategoryReq true "Category data"
// @Success 200 {string} string "message":"Category created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/category/create [post]
func (h *Handlers) CreateCategory(c *gin.Context) {
	var req pb.CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("create-cat", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Category created successfully"})
}

// @Summary Get Category
// @Description Get an Category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {object} pb.Categories "Category data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Category not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/category/{id} [get]
func (h *Handlers) GetCategory(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Category.GetCategory(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Category
// @Description Update an existing Category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Category body pb.UpdateCategoryReq true "Category update data"
// @Success 200 {string} string "message":"Category updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/category/update/{id} [put]
func (h *Handlers) UpdateCategory(c *gin.Context) {
	var req pb.UpdateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("update-cat", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Category updated successfully"})
}

// @Summary List Categorys
// @Description List Categorys with filters
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query int false "Name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListCategoriesRes "List of Categorys"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/category/list [get]
func (h *Handlers) ListCategorys(c *gin.Context) {
	var filter pb.ListAllCategoriesReq

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

	resp, err := h.Category.ListCategories(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Category
// @Description Delete an Category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {string} string "message":"Category deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/category/delete/{id} [delete]
func (h *Handlers) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Category.DeleteCategory(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}
