package controller

import (
	"net/http"
	"time"

	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/dashboardGraphOwl"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/dashboardScreenOWl"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/expression"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/graph"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/host"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/mockcfg"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/strategy"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/template"
	"github.com/Cepave/open-falcon-backend/modules/api/app/controller/uic"
	"github.com/Cepave/open-falcon-backend/modules/api/app/utils"
	"gopkg.in/gin-contrib/cors.v1"
	"gopkg.in/gin-gonic/gin.v1"
)

var headers = []string{
	"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cache-Control", "X-Requested-With",
	"accept", "origin", "Apitoken",
	"page-size", "page-pos", "order-by", "page-ptr", "total-count", "page-more", "previous-page", "next-page",
}

var corsConfig cors.Config

func StartGin(port string, r *gin.Engine) {
	corsConfig = cors.Config{
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowAllOrigins:  true,
		AllowHeaders:     headers,
		ExposeHeaders:    headers,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))
	r.Use(utils.CORS())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm OWL (｡A｡)")
	})
	graph.Routes(r)
	uic.Routes(r)
	template.Routes(r)
	strategy.Routes(r)
	host.Routes(r)
	expression.Routes(r)
	mockcfg.Routes(r)
	dashboardScreenOWl.Routes(r)
	dashboardGraphOwl.Routes(r)
	r.Run(port)
}
