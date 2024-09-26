package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Cart
// @Description Create a new Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Cart body pb.CreateCartReq true "Cart data"
// @Success 200 {string} string "message":"Cart created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/cart/create [post]
func (h *Handlers) CreateCart(c *gin.Context) {
	var req pb.CreateCartReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("create-cart", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Cart created successfully"})
}

// @Summary Get Cart
// @Description Get an Cart by ID
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} pb.Cart "Cart data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Cart not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/cart/{id} [get]
func (h *Handlers) GetCart(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Cart.GetCart(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary List Carts
// @Description List Carts with filters
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart_name query int false "Cart Name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListAllCartRes "List of Carts"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/cart/list [get]
func (h *Handlers) ListAllCarts(c *gin.Context) {
	var filter pb.ListAllCartReq

	cart_name := c.Query("cart_name")
	filter.CartName = cart_name

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

	resp, err := h.Cart.ListAllCarts(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}
