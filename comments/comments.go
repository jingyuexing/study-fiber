/*
* @Author: Jingyuexing
* @Date:   2022-06-29 00:37:11
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-07-03 00:11:23
 */

package comments

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindPost(postid int){

}

func CommentsPostHandler(ctx *fiber.Ctx) error {
  var _postid string = ctx.Params("postid")
  postid,err := strconv.Atoi(_postid)
  if err != nil {
    panic(err)
  }else{
    FindPost(postid)
  }
  return ctx.Next()
}

func CommentsDeleteHandler(ctx *fiber.Ctx) error {
  commentID,err := ctx.ParamsInt("commentID")
  if err != nil {
    return ctx.JSON(fiber.Map{
      "message":"Sorry, something went wrong",
    })
  }
  return ctx.JSON(fiber.Map{
    "comment": commentID,
    "message": "删除评论",
  })
}

func CommentsListHandler(ctx *fiber.Ctx) error {
  postid,err := ctx.ParamsInt("postid")
  if err != nil {
    return ctx.JSON(fiber.Map{
      "message":"Sorry, something went wrong",
    })
  }
  return ctx.JSON(fiber.Map{
    "postid": postid,
    "message":"评论列表",
    "status": 200,
  })

}

