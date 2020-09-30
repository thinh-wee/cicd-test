package template

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sirupsen/logrus"
)

var (
	entry *logrus.Entry
	serv  Service
)

func init() {
	entry = newEntry()
	serv = newService()
}

func newEntry() *logrus.Entry {

	/* CREATE A FORMATTER */
	formatter := new(logrus.TextFormatter)
	formatter.DisableColors = true //false
	formatter.ForceQuote = true
	formatter.FullTimestamp = true
	formatter.PadLevelText = true
	formatter.QuoteEmptyFields = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	newLogger := logrus.New()

	newLogger.SetFormatter(formatter)
	newLogger.SetLevel(logrus.DebugLevel)
	newLogger.SetOutput(os.Stdout)

	return logrus.NewEntry(newLogger)
}

func Test_download(t *testing.T) {
	h := &handler{
		service: serv,

		log: entry,
	}

	// CASE BEAUTIFUL : standard value input
	var (
		queryKey   string
		queryValue []byte
	)
	for k, v := range mockDB {
		queryKey = k
		queryValue = v
	}
	queryID, err := primitive.ObjectIDFromHex(queryKey)
	if err != nil {
		t.Logf("case standard value input failed by : %s [err= %+v]", "bad-error", err)
		t.FailNow()
	}
	statusCode, response, err := downloadExecute(h, &downloadReq{
		ID: queryID,
	})
	if err != nil {
		t.Logf("case standard value input failed by : %s [err= %+v]", "bad-error", err)
		t.FailNow()
	}
	if statusCode != http.StatusOK {
		t.Logf("case standard value input failed by : %s [statusCode= %+v]", "bad-status", statusCode)
		t.FailNow()
	}
	if bytes.Compare(response.Data, queryValue) != 0 {
		t.Logf("case standard value input failed by : %s [response.Data= %+v]", "bad-response", response.Data)
		t.FailNow()
	}

	// CASE 1 : input with ID does not exist
	statusCode, response, err = downloadExecute(h, &downloadReq{
		ID: primitive.NewObjectID(),
	})
	if err != nil {
		t.Logf("case input with ID 'does not exist' failed by : %s [err= %+v]", "bad-error", err)
		t.FailNow()
	}
	if statusCode != http.StatusOK {
		t.Logf("case input with ID 'does not exist' failed by : %s [statusCode= %+v]", "bad-status", statusCode)
		t.FailNow()
	}
	if response.Data != nil {
		t.Logf("case input with ID 'does not exist' failed by : %s [statusCode= %+v]", "bad-response", response.Data)
		t.FailNow()
	}

	// CASE 2 : input with ID invalid
	statusCode, response, err = downloadExecute(h, &downloadReq{})
	if err != nil {
		// pass
	} else {
		t.Logf("case input with ID invalid failed by : %s [err= <nil>]", "bad-error")
		t.FailNow()
	}
	if len(response.Data) != 0 {
		t.Logf("case input with ID invalid failed by : %s [response.Data= %+v]", "bad-response", response.Data)
		t.FailNow()
	}
	if statusCode != http.StatusBadRequest {
		t.Logf("case input with ID invalid failed by : %s [statusCode= %+v]", "bad-status", statusCode)
		t.FailNow()
	}

}

func downloadExecute(h *handler, req *downloadReq) (int, *downloadResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return h.service.download(ctx, req)
}
