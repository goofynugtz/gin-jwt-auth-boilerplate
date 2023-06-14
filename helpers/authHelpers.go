package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, role string) (err error) {
	userType := ctx.GetString("user_type")
	fmt.Println(userType, role)
	err = nil
	if userType != role {
		err = errors.New("unauthorized access")
		return err
	}
	return err
}

func MatchUserTypeToUId(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized access")
		return err
	}

	err = CheckUserType(ctx, userType)
	return err
}
