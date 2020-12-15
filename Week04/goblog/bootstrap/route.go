package bootstrap

import (
	"github.com/bai615/Go-000/Week04/goblog/pkg/route"
	"github.com/bai615/Go-000/Week04/goblog/routes"
	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}