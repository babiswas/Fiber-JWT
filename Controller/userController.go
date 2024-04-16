package Controller

import (
	"newappp/Database"
	"newappp/Model"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *fiber.Ctx) error {

	var body struct {
		UserName string
		Email    string
		Password string
	}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "Invalid Data"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Failed to hash password."})
	}

	user := Model.User{Email: body.Email, UserName: body.UserName, Password: string(hash), CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	result := Database.DB.Create(&user)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Failed to create user."})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User created sucessfully",
	})
}

func Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "Invalid Data."})
	}

	var user Model.User
	Database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "User not found."})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.Status(400).JSON(fiber.Map{"success": false, "message": "Password Mismatch."})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": user.ID, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Failed to create token."})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"token": tokenString,
	})
}

func GetAllUser(ctx *fiber.Ctx) error {
	var users []Model.User
	Database.DB.Find(&users)
	return ctx.Status(200).JSON(fiber.Map{"users": users})
}
