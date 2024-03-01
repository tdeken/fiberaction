package action

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

type Demo struct {
}

func (d Demo) Register() []Action {
	return []Action{
		NewAction(fiber.MethodGet, d.hello),
	}
}

func (d Demo) Group() string {
	return "demo"
}

func (d Demo) ChooseMid(t MidType) []fiber.Handler {
	return []fiber.Handler{func(ctx *fiber.Ctx) error {

		fmt.Println(ctx.BaseURL())

		return nil
	}}
}

func (d Demo) hello(ctx *fiber.Ctx) (err error) {
	return
}

func TestController(t *testing.T) {
	ser := fiber.New()

	AutoRegister(ser, Demo{})

	go func() {
		ser.Listen(":8080")
	}()

	//监听程序退出
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	select {
	case <-ch:
		defer close(ch)
	}

	ser.Shutdown()
}
