package controller

import (
	"fmt"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct{}

type LoginParam struct {
	AccountNum string `json:"account_num" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func (self *LoginController) LoginPost(c *gin.Context) {
	var param LoginParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginUser := model.NewLoginUser(param.AccountNum, param.Password)
	formAccountNum := loginUser.AccountNum
	formPassword := loginUser.Password
	fmt.Println("----------------")
	fmt.Println("[fromAccount]")
	fmt.Println(formAccountNum)
	fmt.Println("[formPassword]")
	fmt.Println(formPassword)
	fmt.Println("----------------")

	dbuser, err := model.GetUserByAccountNum(connecter.DB(), formAccountNum)
	if err != nil {
		// fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	dbAcountNum := dbuser.AccountNum
	dbPassword := dbuser.Password
	fmt.Println("----------------")
	fmt.Println("[dbAcountNum]")
	fmt.Println(dbAcountNum)
	fmt.Println("[dbPassword]")
	fmt.Println(dbPassword)
	fmt.Println("----------------")

	if formAccountNum == dbAcountNum {
		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
			c.JSON(http.StatusOK, gin.H{"結果": "ログイン失敗"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"結果": "ログイン成功"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"結果": "AccountNumが違います"})
}
// 追加メッセージ
// 追加メッセージ2