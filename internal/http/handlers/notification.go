package handlers

import (
	"context"
	"strconv"

	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create a new notification with the provided details
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body pb.NotificationCreate true "Notification details"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/notification/create [post]
func (h *Handlers) CreateNotification(c *gin.Context) {
	var req pb.NotificationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	err = h.Producer.ProduceMessages("notification-create", input)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(201, "Notification created and sent to Kafka successfully")
}

// UpdateNotification godoc
// @Summary Update notification details by ID
// @Description Update the details of a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Param notification body pb.NotificationUpt true "Notification details"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/notification/update/{id} [put]
func (h *Handlers) UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	var body pb.NotificationUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	req := &pb.NotificationUpdate{
		NotificationId: id,
		Body:           &body,
	}

	_, err := h.Notification.UpdateNotificcation(context.Background(), req)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification updated and sent to Kafka successfully")
}

// DeleteNotification godoc
// @Summary Delete a notification by ID
// @Description Delete a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 200 {object} pb.Void
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/notification/delete/{id} [delete]
func (h *Handlers) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetById{Id: id}

	_, err := h.Notification.DeleteNotificcation(context.Background(), req)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification deleted successfully")
}

// GetNotifications godoc
// @Summary List notifications with filters
// @Description Retrieve a list of notifications with optional filters
// @Tags notifications
// @Accept json
// @Produce json
// @Param reciever_id query string false "Receiver ID"
// @Param status      query string false "Notification Status"
// @Param sender_id   query string false "Sender ID"
// @Param limit       query int32 false "Limit for pagination"
// @Param offset      query int32 false "Offset for pagination"
// @Success 200 {object} pb.NotificationList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notifications [get]
func (h *Handlers) ListNotifications(c *gin.Context) {
	var filter pb.NotifFilter
	filter.Filter = &pb.Pagination{}
	filter.RecieverId = c.Query("reciever_id")
	filter.Status = c.Query("status")
	filter.SenderId = c.Query("sender_id")
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Filter.Limit = int32(l)
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Filter.Offset = int32(o)
		}
	}

	resp, err := h.Notification.GetNotifications(context.Background(), &filter)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, resp)
}

// GetNotification godoc
// @Summary Get a notification by ID
// @Description Retrieve a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 200 {object} pb.NotificationGet
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/notification/{id} [get]
func (h *Handlers) GetNotification(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetById{Id: id}

	resp, err := h.Notification.GetNotification(context.Background(), req)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, resp)
}

// NotifyAll godoc
// @Summary Send a notification to all users
// @Description Send a notification message to all users in the specified group
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body pb.NotificationMessage true "Notification message"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/notification/notify-all [post]
func (h *Handlers) NotifyAll(c *gin.Context) {
	var req pb.NotificationMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	_, err := h.Notification.NotifyAll(context.Background(), &req)
	if err != nil {
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification sent to all users successfully")
}
