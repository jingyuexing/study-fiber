package user

import "github.com/gofiber/fiber/v2"

/*
* @Author: Jingyuexing
* @Date:   2022-06-29 00:25:44
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-06-29 02:14:27
 */

func UserSearchHandler(ctx *fiber.Ctx) error {
  uid := ctx.Params("uid")
  name := ctx.Query("name")
  return ctx.JSON(fiber.Map{
    "uid":uid,
    "name":name,
    "message":"search account",
  })
}

func UserRegisterHandler(ctx *fiber.Ctx)error {

  return ctx.JSON(fiber.Map{
    "message":"register account",
  })
}

func UserDeleteHandler(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map{
    "message":"delete account",
  })
}

func UserUpdateHandler(ctx *fiber.Ctx) error {
  uid := ctx.Params("uid")
  return ctx.JSON(fiber.Map{
    "uid": uid,
    "message":"update account",
  })
}

func ResetPasswordHandler(ctx *fiber.Ctx) error{
  return ctx.JSON(fiber.Map{
    "message":"change password",
  })
}

