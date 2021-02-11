package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// TestLoginCreate ...
func TestLoginCreate(t *testing.T) {

	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("ログインが行えること", func(t *testing.T) {
		p := LoginRequest{Email: "hogehoge@gmail.com", Password: "coca cola"}
		byteProduct, _ := json.Marshal(p)
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/v1/login/",
			bytes.NewBuffer(byteProduct),
		)
		assert.Equal(t, 200, response.Code)
		t.Log("--------------")
		t.Log(response.Code)
		t.Log(response.Body)
		t.Log("--------------")
	})
}
