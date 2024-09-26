package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Payment
// @Description Create a new Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Payment body pb.CreatePaymentReq true "Payment data"
// @Success 200 {string} string "message":"Payment created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/payment/create [post]
func (h *Handlers) CreatePayment(c *gin.Context) {
	var req pb.CreatePaymentReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := json.Marshal(req)
	err = h.Producer.ProduceMessages("create-pay", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	c.JSON(200, gin.H{"message": "Payment created successfully"})
}

// @Summary Get Payment
// @Description Get an Payment by ID
// @Tags Payment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} pb.Payment "Payment data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Payment not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/payment/{id} [get]
func (h *Handlers) GetPayment(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Payment.GetPayment(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary List Payments
// @Description List Payments with filters
// @Tags Payment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payment_method query int false "Payment Method"
// @Param status query int false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListPaymentsRes "List of Payments"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/payment/list [get]
func (h *Handlers) ListAllPayments(c *gin.Context) {
	var filter pb.ListPaymentsReq

	payment_metod := c.Query("payment_method")
	status := c.Query("status")
	filter.PaymentMethod = payment_metod
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

	resp, err := h.Payment.ListPayments(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}
