/*
* @Author: Jingyuexing
* @Date:   2022-06-29 00:30:51
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-07-03 00:14:39
 */

package post

import (

	"github.com/gofiber/fiber/v2"
)

func PostCreateHandler(ctx *fiber.Ctx) error {
  return ctx.SendString("new Post")
}

func GetPostHandler(ctx *fiber.Ctx) error {
  postId,err := ctx.ParamsInt("postid")

  if err != nil {
    return ctx.JSON(fiber.Map{
      "id": postId,
      "message":"some thing is wrong",
    })
  }
  return ctx.JSON(fiber.Map{
    "id":postId,
    "message": "你获取到了一个帖子",
  })
}

func UpdatePostHandler(ctx *fiber.Ctx) error {
  postid := ctx.Params("id")
  return ctx.JSON(fiber.Map{
    "id":postid,
    "message": "message information",
  })
}

func PostDeleteHandler(ctx *fiber.Ctx) error {
  postID := ctx.Path("postid")
  return ctx.JSON(fiber.Map{
    "message": "delete post",
    "postid": postID,
  })
}

func PostPutHandler(ctx *fiber.Ctx) error {
  return ctx.SendString("put post")
}
