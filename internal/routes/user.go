package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/handler/http"
	"go-fiber-api/pkg/middleware"
)

func UserV1Route(v1 fiber.Router, hdl *http.Handler) {
	public := v1.Group("/users")
	//private route
	middleware := middleware.NewMiddleware()
	private := middleware.JWTMiddleware(public)
	{
		private.Get("", hdl.GetAllUser)
		private.Get("/:id", hdl.FindUserById)
		//r.Post("", hdl.CreateUser)
		//r.Put("/:id", hdl.UpdateUser)
		//r.Delete("/:id", hdl.DeleteUser)
	}
}
