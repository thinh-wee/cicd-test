package template

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
}

/*
Service repository
*/
type Service interface {
	download(ctx context.Context, request *downloadReq) (statusCode int, response *downloadResp, err error)
}

func newService() Service {
	return &service{}
}

func (ins *service) download(ctx context.Context, request *downloadReq) (statusCode int, response *downloadResp, err error) {
	if err := request.validate(); err != nil {
		return http.StatusBadRequest, &downloadResp{}, err
	}

	// query
	output := func(key primitive.ObjectID) []byte {
		// mocks
		if mockDB != nil {
			return mockDB[key.Hex()]
		}
		return nil
	}(request.ID)

	return http.StatusOK, &downloadResp{
		Data: output,
	}, nil
}
