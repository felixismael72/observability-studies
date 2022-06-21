package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Check struct {
	AppName  string `json:"app_name"`
	Status   string `json:"status"`
	Emoticon string `json:"emoji"`
}

type CheckHandler struct{}

func NewCheckHandler() *CheckHandler {
	return &CheckHandler{}
}

func (h CheckHandler) Check(context echo.Context) error {
	return context.JSON(http.StatusOK, &Check{
		AppName:  "library",
		Status:   "GREEN",
		Emoticon: `ʕ•́ᴥ•̀ʔっ`,
	})
}
