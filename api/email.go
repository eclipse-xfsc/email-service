package api

import (
	"errors"
	"net/http"

	"github.com/eclipse-xfsc/email-service/common"
	"github.com/eclipse-xfsc/email-service/handlers"

	"github.com/gin-gonic/gin"
)

var log = common.GetLogger()

func EmailRoute(r *gin.RouterGroup) *gin.RouterGroup {
	g := r.Group("/email")
	g.POST("/new", func(c *gin.Context) {
		log.Info("New Email Send API request")
		err := handlers.SendEmailNew(c)
		var er common.EmailError
		if err == nil {
			c.JSON(200, gin.H{
				"result": "email sent",
			})
		} else if errors.As(err, &er) {
			c.JSON(er.Code, gin.H{
				"error": er.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	g.POST("/new/queue", func(c *gin.Context) {
		log.Info("New Email Send via NATS API request")
		err := handlers.SendEmailNewViaNats(c)
		var er common.EmailError
		if err == nil {
			c.JSON(200, gin.H{
				"result": "send email request sent to queue",
			})
		} else if errors.As(err, &er) {
			c.JSON(er.Code, gin.H{
				"error": er.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	return g
}
