package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/slh335/hpi-mensa-api/components"
	. "github.com/slh335/hpi-mensa-api/types"
	"github.com/slh335/hpi-mensa-api/util"
)

func (server *Server) Index(c echo.Context) error {
	lang := German
	if c.Path() == "/en" {
		lang = English
	}

	meals, err := server.MealService.Get(Location{Id: 9601}, lang, time.Now())
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "failed to load meals",
		})
	}
	return c.HTML(http.StatusOK, util.RenderComponent(components.Index(lang, meals)))
}
