package action

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
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
	return []fiber.Handler{func(ctx fiber.Ctx) error {

		fmt.Println(ctx.BaseURL())

		return ctx.Next()
	}}
}

func (d Demo) hello(ctx fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}

func TestController(t *testing.T) {
	ser := fiber.New()

	AutoRegister(ser, Demo{})

	resp, err := ser.Test(httptest.NewRequest(fiber.MethodGet, "/demo/hello", nil))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusNoContent {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusNoContent)
	}
}
