package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @Description Create a new Order
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Order body pb.CreateOrderReq true "Order data"
// @Success 200 {string} string "message":"Order created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/create [post]
func (h *Handlers) CreateOrder(c *gin.Context) {
	var req pb.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("create-order", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Order created successfully"})
}

// @Summary Get Order
// @Description Get an Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} pb.Orders "Order data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/{id} [get]
func (h *Handlers) GetOrder(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Order.GetOrder(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Order
// @Description Update an existing Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Order body pb.UpdateOrderReq true "Order update data"
// @Success 200 {string} string "message":"Order updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/update/{id} [put]
func (h *Handlers) UpdateOrder(c *gin.Context) {
	var req pb.UpdateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("update-order", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Order updated successfully"})
}

// @Summary List Orders
// @Description List Orders with filters
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param status query int false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListOrdersRes "List of Orders"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/list [get]
func (h *Handlers) ListOrders(c *gin.Context) {
	var filter pb.ListOrdersReq

	status := c.Query("status")
	filter.Status = status

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

	resp, err := h.Order.ListAllOrders(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Order
// @Description Delete an Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {string} string "message":"Order deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/delete/{id} [delete]
func (h *Handlers) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Order.DeleteOrder(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Order deleted successfully"})
}

// @Summary Get Order By Product ID
// @Description Get an Order by Product ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id path string true "Product ID"
// @Success 200 {object} pb.GetOrdersRes "Order data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/by-product/{id} [get]
func (h *Handlers) GetOrderByProductID(c *gin.Context) {
	productID := c.Param("id")

	if productID == "" {
		c.JSON(400, gin.H{"error": "Product ID cannot be empty"})
		return
	}

	if _, err := uuid.Parse(productID); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Product ID format"})
		return
	}

	req := pb.OrderByProductId{
		ProductID: productID,
	}

	res, err := h.Order.GetOrderByProductID(context.Background(), &req)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "No orders found for the product"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(200, res)
}

