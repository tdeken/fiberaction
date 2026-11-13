package action

import "github.com/gofiber/fiber/v3"

// AutoRegister 自动注册路由
func AutoRegister(router fiber.Router, cs ...Controller) {
	for _, c := range cs {
		g := router.Group(c.Group())
		for _, v := range c.Register() {
			middleware := c.ChooseMid(v.midType)
			handlers := make([]any, 0, len(middleware)+1)
			for _, handler := range middleware {
				handlers = append(handlers, handler)
			}
			handlers = append(handlers, v.do)
			g.Add([]string{v.method}, v.createLastPath(), handlers[0], handlers[1:]...)
		}

	}
}
