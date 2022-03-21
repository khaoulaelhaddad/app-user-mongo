package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"userApp/storage/mock"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []mock.Request{
		{
			Body:         `{}`,
			ExpectedCode: http.StatusOK,
			UrlParams:    nil,
			Action:       "Get",
			Path:         "/api/v1/ping",
		},
	}

	for i, test := range tests {
		recorder, mockContext := mock.GetRecorderContext(test)
		Ping()(mockContext)
		if recorder.Code != test.ExpectedCode {
			t.Errorf("Test case %d: Expected %d, got %d", i, test.ExpectedCode, recorder.Code)
		}
	}
}
