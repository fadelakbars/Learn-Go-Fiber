package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserFindAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserCreate(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user) 
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "Gagal memparsing body request",
			"error" : errValidate.Error(),
		})
	}


	newUsr := entity.User{
		Name:  user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUsr).Error 

	if errCreateUser != nil {
		log.Println(errCreateUser)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Gagal menambahkan user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUsr,
	})

}