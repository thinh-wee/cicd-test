package template

import (
	"net/http"
	"time"

	"github.com/thinh-wee/cicd-test/pkg/net"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service

	log *logrus.Entry
}

/*
Handler repository
*/
type Handler interface {
	AddRouters(r *gin.Engine)
}

/*
NewHandler create a new handler repository
*/
func NewHandler(entry *logrus.Entry, s Service) Handler {
	return &handler{
		service: s,

		log: entry,
	}
}

func (ins *handler) AddRouters(r *gin.Engine) {
	var path = r.Group("/temlate", func(c *gin.Context) {
		ins.log.Warningf("| %v | %15v | %10v | %v", c.Writer.Status(), c.ClientIP(), c.Request.Method, c.Request.RequestURI)
	})

	path.Handle(http.MethodGet, "/download", ins.download)
}

func (ins *handler) download(c *gin.Context) {
	defer func(begin time.Time) {
		ins.log.Warningf("Processing ended in %v \r\n ------------------------------", time.Since(begin))
	}(time.Now())

	var (
		request  = &downloadReq{}
		response = &downloadResp{}

		statusCode int

		r = c.Request
		w = c.Writer
	)

	net.BindJSON(ins.log, r, request)

	statusCode, response, _ = ins.service.download(r.Context(), request)

	net.WriteRaw(ins.log, r, w, statusCode, nil, response.Data)
}
