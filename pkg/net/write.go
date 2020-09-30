package net

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
WriteRaw func
*/
func WriteRaw(log *logrus.Entry, r *http.Request, w gin.ResponseWriter, statusCode int, header map[string][]string, raw []byte) (err error) {
	defer func() {
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"body": base64.StdEncoding.EncodeToString(raw), // fmt.Sprintf("%s", raw),
			"err":  err,
		}).Debugf("")
		// logrus.Debugf("[WriteJSON] --- %v --- %+v", r.RequestURI, len(raw))
		if err != nil {
			logrus.Error(err)
		}
	}()
	for k, h := range header {
		for _, v := range h {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(raw)
	return err
}

/*
WriteJSON func
*/
func WriteJSON(log *logrus.Entry, r *http.Request, w gin.ResponseWriter, statusCode int, v interface{}) (err error) {
	defer func() {
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"json": fmt.Sprintf("%+v", v),
			"err":  err,
		}).Debugf("")
		// logrus.Debugf("[WriteJSON] --- %v --- %+v", r.RequestURI, v)
		if err != nil {
			logrus.Error(err)
		}
	}()
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(v)
}

/*
WriteHTML func
*/
func WriteHTML(log *logrus.Entry, r *http.Request, w gin.ResponseWriter, statusCode int, textHTML string) (err error) {
	defer func() {
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"body": fmt.Sprintf("%s", textHTML),
			"err":  err,
		}).Debugf("")
		if err != nil {
			logrus.Error(err)
		}
	}()
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, err = w.WriteString(textHTML)
	return err
}

/*
WriteError func
*/
func WriteError(log *logrus.Entry, r *http.Request, w gin.ResponseWriter, statusCode int, err error) (writeErr error) {
	var v = struct {
		IsError bool
		Message string
	}{
		IsError: true,
		Message: err.Error(),
	}
	defer func() {
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"json": fmt.Sprintf("%+v", err),
			"err":  writeErr,
		}).Debugf("")
		// logrus.Debugf("[WriteError] -- %v --- %v", r.RequestURI, err.Error())
		if writeErr != nil {
			logrus.Error(writeErr)
		}
	}()
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	return json.NewEncoder(w).Encode(v)
}
