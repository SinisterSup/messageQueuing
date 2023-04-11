package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SinisterSup/messageQueuing/internal/app/api"
	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api.SetupRoutes(router)

	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/invalid-route", nil)
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected status code %v, but got %v", http.StatusNotFound, resp.Code)
	}
}
