package main

import (
	"DATABASE_MONEYTAPS/models"
	"DATABASE_MONEYTAPS/storage"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Register(context *fiber.Ctx) error {
	user := models.User{}

	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Check if the user with the same contact information already exists
	existingUser := models.User{}
	err = r.DB.Where("contact = ?", user.Contact).First(&existingUser).Error
	if err == nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "User with the same contact already exists"})
		return err
	}

	// Create the user
	err = r.DB.Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not register the user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Successfully Registered",
	})
	return nil
}

/*func (r *Repository) Register(context *fiber.Ctx) error {
user := models.User{}

 err := context.BodyParser(&user)
 if err != nil {
 context.Status(http.StatusUnprocessableEntity).JSON(
 &fiber.Map{"message": "request failed"})
 return err
 }

 err = r.DB.Create(&user).Error
 context.Status(http.StatusBadRequest).JSON(
 &fiber.Map{"message": "could not register the user"})
}

context.Status(http.StatusOK).JSON(&fiber.Map{
 "message": "Successfully Registered",
 })
 return nil
}*/

func GenerateJWT(userID uint) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		// Add other desired claims...
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the token string
	jwtSecret := []byte("your_secret_key") // Replace with your secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (r *Repository) Login(context *fiber.Ctx) error {
	user := models.User{}

	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Perform authentication logic
	// Example:
	userModel := &models.User{}
	err = r.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(userModel).Error
	if err != nil {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "invalid email or password"})
		return err
	}

	// Assuming authentication is successful, generate a token
	// Example:
	token, err := GenerateJWT(userModel.ID)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to generate token"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login successful",
		"token":   token,
	})
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
	err := r.DB.Delete(userModel, id)

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

	err := r.DB.Find(userModel).Error
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

	err := r.DB.Where("id = ?", id).First(userModel).Error
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
func (r *Repository) CreateLoan(context *fiber.Ctx) error {
	loan_statuses := models.LoanStatus{}
	status := "0"

	err := context.BodyParser(&loan_statuses)
	fmt.Println(loan_statuses)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	loan_statuses.Status = status
	err = r.DB.Create(&loan_statuses).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create client"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "client has been added"})
	return nil
}

func (r *Repository) GetAllLoan(context *fiber.Ctx) error {
	loan_statuses := &[]models.LoanStatus{}

	err := r.DB.Find(loan_statuses).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user data",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "user successfully fetch", "data": loan_statuses})
	return nil
}
func (r *Repository) ApprovedLoan(context *fiber.Ctx) error {
	loan_statuses := models.LoanStatus{}

	err := context.BodyParser(&loan_statuses)
	fmt.Println(loan_statuses)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&loan_statuses).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not been approved"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "your load has been Approved"})
	return nil
}
func (r *Repository) ClientHistory(context *fiber.Ctx) error {
	history := models.History{}

	err := context.BodyParser(&history)
	fmt.Println(history)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&history).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create history client"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "history client has been added"})
	return nil
}
func (r *Repository) DeleteLoanStatus(context *fiber.Ctx) error {
	loanModel := models.LoanStatus{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := r.DB.Delete(loanModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete loan",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "loan deleted successfully",
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_user", r.Register)
	api.Post("/user/login", r.Login) // Updated route for user login
	api.Delete("delete_user/:id", r.DeleteUser)
	api.Get("/get_user/:id", r.GetUserByID)
	api.Get("/user", r.GetUser)
	api.Post("/create_loan", r.CreateLoan)
	api.Get("/getall_loanstatus", r.GetAllLoan)
	api.Delete("delete_loan_stastus/:id", r.DeleteLoanStatus)
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
