package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestHealthCheck(t *testing.T) {

	t.Run("200を返却すること", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := Router()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

}
