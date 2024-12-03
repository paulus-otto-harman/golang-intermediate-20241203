package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/class/database"
	"project/class/handler"
)

type Middleware struct {
	Cacher database.Cacher
}

func NewMiddleware(cacher database.Cacher) Middleware {
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("ID-KEY")
		log.Println(token, id)
		val, err := m.Cacher.Get(id)
		if err != nil {
			handler.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if val == "" || val != token {
			handler.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		log.Println("pass middleware")

		// before request
		c.Next()

	}
}
