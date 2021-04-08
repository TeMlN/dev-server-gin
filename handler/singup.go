package handler

import "github.com/gin-gonic/gin"

type signReq struct {
	Id    string `json:"id" query:"id"`
	Email string `json:"email" query:"id"`
	Pwd   string `json:"pwd" query:"pwd"`
}

func SignUp(c *gin.Context) {}
	S := new(signReq) 
}
