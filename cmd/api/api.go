package api

import (
	"context"
	"database/sql"
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

func (s *APIServer) Run() error {
	router := GetRouter()
	
	return http.ListenAndServe(s.Addr, router)
}


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
		OperationID:   " helloworld",
		Method:        http.MethodPost,
		Path:          apiPath + "/helloworld",
		Summary:       "",
		Description:   "HelloWorldbasic",
		Tags:          []string{},
		Errors:        []int{400},
		DefaultStatus: 201,
	}, func(ctx context.Context, input *struct{}) (*struct{}, error) {
		return nil, nil
	})
	
	huma.Register(routerReadApi, huma.Operation{
		OperationID:   "getstarted",
		Method:        http.MethodGet,
		Path:          apiPath + "/Gettingstarted",
		Summary:       "setting up basics",
		Description:   "This endpoint returns mechanical component serial number metadata.",
		Tags:          []string{"setting basics"},
		Errors:        []int{400},
		DefaultStatus: 200,
	}, func(ctx context.Context, input *struct{}) (*struct{}, error) {
		return nil, nil
	})

return router 
}

