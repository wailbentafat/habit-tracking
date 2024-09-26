package controlers

import (
	"habit/db"
	"habit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProgressAnalytics(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var habits []models.Habit
	if err := db.DB.Where("userid = ?", userID).Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching habits"})
		return
	}

	var analytics []gin.H

	for _, habit := range habits {
		var progresses []models.Progres
		if err := db.DB.Where("habit_id = ?", habit.ID).Find(&progresses).Error; err != nil {
			continue 
		}

		total := len(progresses)
		completed := 0
		missed := 0

		for _, progress := range progresses {
			if progress.Status == "completed" {
				completed++
			} else if progress.Status == "missed" {
				missed++
			}
		}

		analytics = append(analytics, gin.H{
			"habit_id":   habit.ID,
			"habit_name": habit.Name,
			"total":      total,
			"completed":  completed,
			"missed":     missed,
			"completion_rate": float64(completed) / float64(total) * 100,
		})
	}

	c.JSON(http.StatusOK, gin.H{"analytics": analytics})
}
