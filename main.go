package main


import(
	"github.com/gofiber/fiber/v2"
)

func main(){
	route:=fiber.New()
	route.Get("/",func (ctx *fiber.Ctx)error  {
		return ctx.SendString("Hello world")
	})
	route.Listen(":80")
}
