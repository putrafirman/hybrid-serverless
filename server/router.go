package server

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"putrafirman.com/playground/hybrid-serverless/handler"
)

func RegisterRouter(e *echo.Echo) *echo.Echo {

	// DB Conn
	db := dbConn()
	e.Logger.Info(db.Ping())
	defer db.Close()

	pingHandler := handler.NewPingHandler(db)

	// Routes
	e.GET("/", hello)
	e.GET("/ping", pingHandler.GetPing())

	return e
}

// Example handler without layering architecture
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World from gin!")
}

// TODO - separate init connection for DB
func dbConn() (db *sql.DB) {
	// TODO - must be in env var for production use. for intial will be hardcoded
	dbDriver := "mysql"
	dbUser := ""
	dbPass := ""
	dbName := ""
	dbHost := ""

	dbParam := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open(dbDriver, dbParam)
	if err != nil {
		panic(err.Error())
	}
	return db
}
