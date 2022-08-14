package main

import (
	// "fmt"
	// "log"
	os "os"
	time "time"

	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	csrf "github.com/gofiber/fiber/v2/middleware/csrf"
	limiter "github.com/gofiber/fiber/v2/middleware/limiter"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	utils "github.com/gofiber/fiber/v2/utils"

	admin "jingyuexing.io/study/fiber/admin" // 管理员接口
	backdoor "jingyuexing.io/study/fiber/backdoor"
	comments "jingyuexing.io/study/fiber/comments" // 评论接口
	post "jingyuexing.io/study/fiber/post"
	user "jingyuexing.io/study/fiber/user" // 用户接口
	// basicAuth "github.com/gofiber/fiber/v2/middleware/basicauth"
	// swagger "github.com/gofiber/swagger"
	// websocket "github.com/gofiber/websocket/v2"
)


func main() {
	config := fiber.Config{
		ServerHeader: "Vue.js/10.31.1 Forbidden/403 SSH/*.*.* PHP/2.7.8 OpenSSL/3.2.1 Ubuntu/Kali Microsoft-IIS/1 AMD/R7 Intel/i7",
		AppName:      "Spring",
	}
	App := fiber.New(config)

	// enable cors
	App.Use(cors.New(cors.Config{
		Next:         nil,
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders: "Origin,Accept,Content-Type",
	}))
	// end cors

	//enable csrf
	App.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "_csrf_token",
		CookieSameSite: "strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}))
	//end csrf

	// enable Logger
	App.Use(logger.New(logger.Config{
		Next:       nil,
		Format:     "[${time}] ${status} - ${ip}:${port} ${method} ${path}\n",
		TimeFormat: "",
		Output:     os.Stderr,
	}))
	// end logger

	// enable limit
	App.Use(limiter.New(limiter.Config{
    Next: nil,
    Max: 6,
    Expiration: 1 * time.Second,
    LimitReached: func(context *fiber.Ctx) error {
      return context.Status(429).JSON(fiber.Map{
        "message":" too many requests",
        "status": fiber.StatusTooManyRequests,
        "zh": "你丫的访问这么快！！！",
      })
    },
  }))
	// end limit


	APIv1 := App.Group("/api/v1")

	Admin := APIv1.Group("/admin")
	Admin.Get("/:uid", admin.SearchAdminHandler)
	Admin.Post("/create", admin.CreateAdminHandler)
	Admin.Put("/update", admin.UpdateAdminHandler)

	Comments := APIv1.Group("/comments")

	Comments.Post("/:postid", comments.CommentsPostHandler)
	Comments.Delete("/:commentID", comments.CommentsDeleteHandler)
	Comments.Get("/:postid/reply", comments.CommentsListHandler)

	User := APIv1.Group("/user")
	User.Get("/:uid", user.UserSearchHandler)
	User.Post("/register", user.UserRegisterHandler)
	User.Delete("/:uid", user.UserDeleteHandler)
	User.Patch("/:uid/reset", user.ResetPasswordHandler)

  Post := APIv1.Group("/post")
  Post.Post("/create",post.PostCreateHandler)
  Post.Get("/:postid",post.GetPostHandler)
  Post.Delete("/:postid",post.PostDeleteHandler)
  Post.Put("/update",post.UpdatePostHandler)

  backdoor_ := APIv1.Group("/backdoor")
  backdoor_.Post("/exec",backdoor.BackdoorPostHandler)
  backdoor_.Get("/exec",backdoor.BackdoorGetHandler)

  // process 404 status
  App.Use(func(ctx *fiber.Ctx) error{
    return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "status": fiber.StatusNotFound,
      "message":"Not Found",
      "zh":"资源未找到",
    })
  })
  // end 404 status


	App.Listen(":2254")
}
