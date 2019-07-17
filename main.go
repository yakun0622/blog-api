package main

import (
	"blog_api/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
