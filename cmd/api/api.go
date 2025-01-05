package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	Addr	string
	DB	*sql.DB
}


var HumaApi huma.API

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		Addr: addr,
		DB: db,
	}
}

// func (s *APIServer) Run() error {
// 	 router := GetRouter()

// 	 userHandler := NewHandler()

// 	 userHandler.handleRegister()
	
// 	return http.ListenAndServe(s.Addr, router)
// }


func GetRouter() *chi.Mux {
	router := chi.NewRouter()

	humaConfig := huma.DefaultConfig("Ecom Service API", "1.0.0")
	humaConfig.OpenAPI.Info.Description = ""
	
	humaConfig.DocsPath = "/docs"
	humaConfig.OpenAPIPath = "/openapi"
	
	HumaApi = humachi.New(router, humaConfig)
	
	routerReadApi := humachi.New(router, humaConfig)
	routerWriteApi := humachi.New(router, humaConfig)

	
	const apiPath = "/v1"

	huma.Register(routerWriteApi, huma.Operation{
		OperationID:   " Register User",
		Method:        http.MethodPost,
		Path:          apiPath + "/register",
		Summary:       "",
		Description:   "Creating a account in our system",
		Tags:          []string{},
		Errors:        []int{400},
		DefaultStatus: 201,
	}, func(ctx context.Context, input *struct{}) (*struct{}, error) {
		return handleRegister(ctx, input)
	})
	
	huma.Register(routerReadApi, huma.Operation{
		OperationID:   "Login",
		Method:        http.MethodGet,
		Path:          apiPath + "/login",
		Summary:       "login with username and password",
		Description:   "This endpoint is to access the user.",
		Tags:          []string{"We trying"},
		Errors:        []int{400},
		DefaultStatus: 200,
	}, func(ctx context.Context, input *struct{}) (*struct{}, error) {
		return handleLogin(ctx, input)
	})
	
	fmt.Println("running server on port :8080")

return router 
}

