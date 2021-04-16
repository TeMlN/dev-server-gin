package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SignReq struct {
	Id    string `json:"id" query:"id" form:"email"`
	Email string `json:"email" query:"email" form:"email"`
	Pwd   string `json:"pwd" query:"pwd" form:"email"`
}

func SignUp(c *gin.Context) {
	s := new(SignReq)

	if err := c.Bind(s); err != nil {
		return
	}

	fmt.Println(s.Id, s.Email, s.Pwd)

	if s.Id == "" || s.Email == "" || s.Pwd == "" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "모든값을 입력해주세요.",
		})
	}

}
