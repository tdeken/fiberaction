package action

import "github.com/gofiber/fiber/v2"

// AutoRegister 自动注册路由
func AutoRegister(router fiber.Router, cs ...Controller) {
	for _, c := range cs {
		g := router.Group(c.Group())
		for _, v := range c.Register() {
			do := c.ChooseMid(v.midType)
			do = append(do, v.do)
			g.Add(v.method, v.createLastPath(), do...)
		}

	}
}
