package api

import (
	"context"
	"fmt"
)


type InputHandler struct {

}



func  handleLogin(ctx context.Context, input *struct{})(*struct{}, error){
	fmt.Println(ctx, input)
	return nil, nil
}


func handleRegister(ctx context.Context, input *struct{})(*struct{}, error){
	fmt.Println(ctx, input)
	return nil, nil
}