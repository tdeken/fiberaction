package action

import "github.com/gofiber/fiber/v3"

type MidContainer struct {
}

type Controller interface {
	Register() []Action                  //需要中间件的方法
	Group() string                       //控制器分组
	ChooseMid(t MidType) []fiber.Handler //选择中间件
}
