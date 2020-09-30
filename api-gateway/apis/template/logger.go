package template

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
	next Service

	log *logrus.Entry
}

/*
NewServiceAsLogger create a new service repository as logger
*/
func NewServiceAsLogger(entry *logrus.Entry) Service {
	return &logger{
		next: newService(),

		log: entry,
	}
}

func (ins *logger) download(ctx context.Context, request *downloadReq) (statusCode int, response *downloadResp, err error) {
	defer func(begin time.Time) {
		ins.log.WithFields(
			logrus.Fields{
				"Method": "Request",
				"json":   fmt.Sprintf("%+v", request),
			},
		).Infof("process_time = %+v", time.Since(begin))
		ins.log.WithFields(
			logrus.Fields{
				"Method":     "Response",
				"StatusCode": statusCode,
				"json":       fmt.Sprintf("%+v", response),
			},
		).Infof("process_time = %+v", time.Since(begin))
		if err != nil {
			ins.log.Error(fmt.Sprintf("[ERR] func(apis/template/service.download) : %+v", err))
		}
	}(time.Now())
	return ins.next.download(ctx, request)
}
