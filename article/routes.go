// routes.go here we ave all the routes for the app

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/article/view/:id", showSingleArticle)
}
