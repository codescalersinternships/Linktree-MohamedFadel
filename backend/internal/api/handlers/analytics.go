package handlers

import (
	"linktree-mohamedfadel-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	AnalyticsService *services.AnalyticsService
}

func NewAnalyticsHandler(analyticsService *services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{AnalyticsService: analyticsService}
}

func (h *AnalyticsHandler) TrackLinkClickHandler(c *gin.Context) {
	id := c.Param("id")
	linkId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	visitorUsername, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.AnalyticsService.TrackLinkClicks(uint64(linkId), visitorUsername.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Click tracked successfully"})
}
