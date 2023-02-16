package handler

import (
	"database/sql"

	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

type pingHandler struct {
	db *sql.DB
}

type PingHandler interface {
	GetPing() echo.HandlerFunc
}

func NewPingHandler(db *sql.DB) PingHandler {
	return &pingHandler{
		db: db,
	}
}

func (p *pingHandler) GetPing() echo.HandlerFunc {
	return func(c echo.Context) error {
		// result, err := repository.GetAll(db)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, nil)
		// }
		return c.JSON(http.StatusOK, "PONG")
	}
}
