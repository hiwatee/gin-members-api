package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Arrange ---
	// p := models.User{
	// 	ID:       1,
	// 	Email:    "test@example.com",
	// 	Password: "password",
	// }
	// p.CreateModel()
	// byteProduct, _ := json.Marshal(p)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/api/v1/login",
		nil,
		// bytes.NewBuffer(byteProduct),
	)

	t.Log(response.Code)

	assert.EqualValues(t, http.StatusOK, response.Code)

}
