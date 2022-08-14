package admin

import "github.com/gofiber/fiber/v2"

/*
* @Author: Jingyuexing
* @Date:   2022-06-29 00:31:23
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-07-03 00:02:09
 */

func CreateAdminHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "create admin",
	})
}

func SearchAdminHandler(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")
	return ctx.JSON(fiber.Map{
		"uid":     uid,
		"message": "search admin account",
	})
}

func UpdateAdminHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "update admin account infor",
	})
}
