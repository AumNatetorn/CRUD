package crud

import (
	"CRUD/app"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Customers struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Age  int    `json:"age" gorm:"column:age"`
}

func (r *Customers) TableName() string {
	return "profile"
}

type Response struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       *Customers `json:"data,omitempty"`
}

type CustomerStorage interface {
	InsertOne(ctx context.Context, req Customers) error
	UpdateOne(ctx context.Context, req Customers) error
	DeleteOne(ctx context.Context, id int) error
	FindOne(ctx context.Context, id int) (ent *Customers, err error)
}

type Handler struct {
	crud CustomerStorage
}

func NewHandler(crud CustomerStorage) *Handler {
	return &Handler{crud: crud}
}

func (h *Handler) CreateCustomerHandler(c *fiber.Ctx) error {
	req := Customers{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Bad Request: %s", err.Error()),
		})
	}

	if err := h.crud.InsertOne(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to create customer: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		StatusCode: fiber.StatusOK,
		Message:    "Create Customer Success",
		Data:       nil,
	})
}

func (h *Handler) UpdateCustomerHandler(c *fiber.Ctx) error {
	req := Customers{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Bad Request: %s", err.Error()),
		})
	}

	if err := h.crud.UpdateOne(c.Context(), req); err != nil {
		if errors.Is(err, app.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(Response{
				StatusCode: fiber.StatusNotFound,
				Message:    fmt.Sprintf("Failed to Update Customer: %s", err.Error()),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Update Customer: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		StatusCode: fiber.StatusOK,
		Message:    "Update Customer Success",
		Data:       nil,
	})
}

func (h *Handler) DeleteCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Err Converse ID: %v", err)
		return err
	}

	if err := h.crud.DeleteOne(c.Context(), idInt); err != nil {
		// Return a BadRequest response if an error occurs
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Delete Customer: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		StatusCode: fiber.StatusOK,
		Message:    "Delete Customer Success",
		Data:       nil,
	})
}

func (h *Handler) GetCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Err Converse ID: %v", err)
		return err
	}

	resp, err := h.crud.FindOne(c.Context(), idInt)
	if err != nil {
		if errors.Is(err, app.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(Response{
				StatusCode: fiber.StatusNotFound,
				Message:    fmt.Sprintf("Failed to Get Customer: %s", err.Error()),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Get Customer: %s", err.Error()),
		})

	}

	return c.Status(fiber.StatusOK).JSON(Response{
		StatusCode: fiber.StatusOK,
		Message:    "Get Customer Success",
		Data: &Customers{
			ID:   resp.ID,
			Name: resp.Name,
			Age:  resp.Age,
		},
	})
}
