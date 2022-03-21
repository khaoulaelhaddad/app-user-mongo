package mock

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type Request struct {
	Body         string
	ExpectedCode int
	UrlParams    []gin.Param
	Action       string
	Path         string
	Description  string
}

func GetRecorderContext(request Request) (*httptest.ResponseRecorder, *gin.Context) {

	recorder := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(recorder)

	mockContext.Params = request.UrlParams

	req, _ := http.NewRequest(request.Action, request.Path,
		bytes.NewBuffer([]byte(request.Body)))

	ctx := context.Background()
	req = req.WithContext(ctx)

	mockContext.Request = req

	return recorder, mockContext
}

func Params(params ...gin.Param) []gin.Param {
	return params
}
