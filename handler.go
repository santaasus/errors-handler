package errorshandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	errors := c.Errors
	if len(errors) > 0 {
		err, ok := errors[0].Err.(*AppError)
		if ok {
			message := MessageResponse{err.Error()}
			switch err.Type {
			case Unauthorized:
				c.JSON(http.StatusUnauthorized, message)
			case ValidationError:
				c.JSON(http.StatusBadRequest, message)
			case InternalServerError:
				c.JSON(http.StatusInternalServerError, message)
			case NotFound:
				c.JSON(http.StatusNotFound, message)
			default:
				c.JSON(http.StatusInternalServerError, MessageResponse{"Sorry, this case in development"})
			}

			return
		}

		c.JSON(http.StatusInternalServerError, MessageResponse{"Sorry, this case in development"})
	}
}
