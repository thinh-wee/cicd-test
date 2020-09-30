package net

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

/*
BindQueryDefault func
*/
func BindQueryDefault(log *logrus.Entry, r *http.Request, key, defaultValue string) (query string) {
	defer func() {
		go log.WithFields(logrus.Fields{
			"from":        r.RequestURI,
			"query_key":   key,
			"query_value": query,
		}).Debugf("")
		// logrus.Debugf("[BindQueryDefault] --- %v --- %+v", r.RequestURI, query)
	}()
	value := r.URL.Query()
	if query = value.Get(key); query != "" {
		return
	}
	return defaultValue
}

/*
BindHeader func
*/
func BindHeader(log *logrus.Entry, r *http.Request, key string) (value string) {
	defer func() {
		go log.WithFields(logrus.Fields{
			"from":         r.RequestURI,
			"header_key":   key,
			"header_value": value,
		}).Debugf("")
		// logrus.Debugf("[BindHeader] --- %v --- %+v", r.RequestURI, value)
	}()
	return r.Header.Get(key)
}

/*
BindRawData func
*/
func BindRawData(log *logrus.Entry, r *http.Request) (body []byte, err error) {
	defer func() {
		r.Body.Close()
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"body": base64.StdEncoding.EncodeToString(body), // fmt.Sprintf("%+v", body),
			"err":  err,
		}).Debugf("")
		// logrus.Debugf("[BindRawData] --- %v --- %+v", r.RequestURI, body)
		if err != nil {
			logrus.Error(err)
		}
	}()
	return ioutil.ReadAll(r.Body)
}

/*
BindJSON func
*/
func BindJSON(log *logrus.Entry, r *http.Request, v interface{}) (err error) {
	var body []byte
	defer func() {
		r.Body.Close()
		go log.WithFields(logrus.Fields{
			"from": r.RequestURI,
			"json": fmt.Sprintf("%+v", v),
			"err":  err,
		}).Debugf("")
		// logrus.Debugf("[BindJSON] --- %v --- %+v", r.RequestURI, v)
		if err != nil {
			logrus.Error(err)
			fmt.Printf("======================================\r\n\tBIND JSON ERROR:\r\n%+v\r\n======================================\r\n", string(body))
		}
	}()
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(body)).Decode(v)
}
