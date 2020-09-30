package template

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* === mock data for test === */

var mockDB = map[string][]byte{
	"5f740b4b4d94ee6fe374033d": []byte("mock-5f740b4b4d94ee6fe374033d"),
	"5f740b50d0cf3998734f1bcc": []byte("mock-5f740b50d0cf3998734f1bcc"),
	"5f740b5ea97cd0acfaa10498": []byte("mock-5f740b5ea97cd0acfaa10498"),
}

/* === basic constants === */

const (
	codeTemplate int = 9999
)

/* === error messages === */

const msgErrTemplate string = "this is template message error"

// IsErrTemplate func validate
func IsErrTemplate(err error) bool {
	if err != nil && err.Error() == msgErrTemplate {
		return true
	}
	return false
}

const msgErrIDInvalid string = "ID is invalid"

// IsErrIDInvalid func validate
func IsErrIDInvalid(err error) bool {
	if err != nil && err.Error() == msgErrIDInvalid {
		return true
	}
	return false
}

/* === schemas === */

type templateSchema struct {
	Text string `json:"text"`
}

/* === object i/o === */

type templateObj struct {
	Schema *templateSchema
}

type downloadReq struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

func (ins *downloadReq) validate() error {
	if ins.ID.IsZero() {
		return errors.New("ID is invalid")
	}
	return nil
}

type downloadResp struct {
	Data []byte
}
