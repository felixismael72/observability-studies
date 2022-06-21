package main

import "github.com/labstack/echo/v4"

func LoadRoutes() *echo.Echo {
	router := echo.New()

	api := router.Group("/api")
	loadAuthorRoutes(api)
	loadCheckRoutes(api)

	return router
}

func mountAuthorHandler() *AuthorHandler {
	authorRepo := NewAuthorRepository()
	authorService := NewAuthorService(authorRepo)
	return NewAuthorHandler(authorService)
}

func loadAuthorRoutes(group *echo.Group) {
	authorGroup := group.Group("/author")

	authorHandler := mountAuthorHandler()

	authorGroup.POST("/new", authorHandler.Post)
}

func mountCheckHandler() *CheckHandler {
	return NewCheckHandler()
}

func loadCheckRoutes(group *echo.Group) {
	checkGroup := group.Group("/check")

	checkHandler := mountCheckHandler()

	checkGroup.GET("", checkHandler.Check)
}
