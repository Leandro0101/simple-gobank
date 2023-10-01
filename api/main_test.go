package api

import (
	"database/sql"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
