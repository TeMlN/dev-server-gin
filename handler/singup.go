package handler

import (
	"github.com/gin-gonic/gin"
)

type SignReq struct {
	Id    string `json:"id" query:"id" form:"email"`
	Email string `json:"email" query:"email" form:"email"`
	Pwd   string `json:"pwd" query:"pwd" form:"email"`
}

func SignUp(c *gin.Context) {
	s := new(SignReq)

}
