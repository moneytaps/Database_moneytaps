package main

import (
	"database/models"
	"database/storage"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

type Client struct {
	First_Name  string `json:"firstname"`
	Middle_Name string `json:"middlename"`
	Sur_Name    string `json:"surname"`
	Birthdate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Primary     string `json:"primary "`
	LoanAmount  int    `json:"loanAmount"`
	LoanDays    int    `json:"loanDays"`
	Interest    int    `json:"interest"`
	TotalAmount int    `json:"totalAmount"`
	Purpose     string `json:"purpose"`
}

type User struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Contact  *string `json:"contact"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateUser(context *fiber.Ctx) error {
	user := User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Table("user").Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create user"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been added"})
	return nil
}

func (r *Repository) DeleteUser(context *fiber.Ctx) error {
	userModel :=
		models.User{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Table("user").Delete(userModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete user",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user deleted successfully",
	})
	return nil
}

func (r *Repository) GetUser(context *fiber.Ctx) error {
	userModel := &[]models.User{}

	err := r.DB.Table("user").Find(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user data",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "user successfully fetch", "data": userModel})
	return nil
}

func (r *Repository) GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	userModel := &models.User{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot found",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Table("user").Where("id = ?", id).First(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user id fetched successfully",
		"data":    userModel,
	})
	return nil
}
func (r *Repository) CreateClient(context *fiber.Ctx) error {
	client := Client{}

	err := context.BodyParser(&client)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Debug().Table("client").Omit("First_Name", "Middle_Name", "Sur_name", "Birthdate", "Gender", "Address", "Primary", "LoanAmount", "LoanDays", "Interest", "TotalAmount", "Purpose").Create(&client).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create client"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "client  has been added"})
	return nil
}
func (r *Repository) GetClient(context *fiber.Ctx) error {
	clientModel := &[]models.Client{}

	err := r.DB.Table("client").Find(clientModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user data",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "user successfully fetch", "data": clientModel})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_user", r.CreateUser)
	api.Delete("delete_user/:id", r.DeleteUser)
	api.Get("/get_user/:id", r.GetUserByID)
	api.Get("/user", r.GetUser)
	api.Post("/create_client", r.CreateClient)
	api.Get("/get_client", r.GetClient)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
