package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"subscriptions-service/internal/model"
	"subscriptions-service/internal/repository"
)

type Handler struct {
	Repo *repository.SubscriptionRepo
}

func NewHandler(repo *repository.SubscriptionRepo) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateSubscription(c *gin.Context) {

	var s model.Subscription

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	s.ID = uuid.New()

	err := h.Repo.Create(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h *Handler) GetAllSubscriptions(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	minPrice, _ := strconv.Atoi(c.Query("min_price"))
	maxPrice, _ := strconv.Atoi(c.Query("max_price"))
	sort := c.Query("sort") // price_asc, price_desc, start_date_asc, start_date_desc

	var (
		subs []model.Subscription
		err  error
	)

	if userID != "" || serviceName != "" || minPrice > 0 || maxPrice > 0 || sort != "" {
		subs, err = h.Repo.GetAllFiltered(userID, serviceName, minPrice, maxPrice, sort)
	} else {
		subs, err = h.Repo.GetAll()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}
func (h *Handler) GetSubscriptionByID(c *gin.Context) {
	id := c.Param("id")
	sub, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}
	c.JSON(http.StatusOK, sub)
}
func (h *Handler) UpdateSubscription(c *gin.Context) {
	id := c.Param("id")
	var s model.Subscription
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Repo.Update(id, s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h *Handler) DeleteSubscription(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
func (h *Handler) GetTotal(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start") // формат YYYY-MM
	end := c.Query("end")     // формат YYYY-MM

	total, err := h.Repo.GetTotal(userID, serviceName, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total})
}
