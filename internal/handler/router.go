package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/subscriptions", h.GetAllSubscriptions)
	r.POST("/subscriptions", h.CreateSubscription)
	r.GET("/subscriptions/:id", h.GetSubscriptionByID)
	r.PUT("/subscriptions/:id", h.UpdateSubscription)
	r.DELETE("/subscriptions/:id", h.DeleteSubscription)
	r.GET("/subscriptions/total", h.GetTotal)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
