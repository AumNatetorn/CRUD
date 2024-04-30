package crud

import (
	"CRUD/app"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	reqBody = `{"id":1,"name":"John Doe","age":13}`
)

func TestCreateHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	app := fiber.New()

	app.Post("/customers", handler.CreateCustomerHandler)

	req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusOK, Message: "Create Customer Success"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestCreateHandler_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(errors.New("duplicate")).Times(1)

	app := fiber.New()

	app.Post("/customers", handler.CreateCustomerHandler)

	req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusBadRequest, Message: "Failed to create customer: duplicate"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestUpdateHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	app := fiber.New()

	app.Put("/customers", handler.UpdateCustomerHandler)

	req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusOK, Message: "Update Customer Success"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestUpdateHandler_RecordNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Return(app.ErrRecordNotFound).Times(1)

	app := fiber.New()

	app.Put("/customers", handler.UpdateCustomerHandler)

	req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusNotFound, Message: "Failed to Update Customer: record not found"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestUpdateHandler_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Return(errors.New("bad request")).Times(1)

	app := fiber.New()

	app.Put("/customers", handler.UpdateCustomerHandler)

	req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusBadRequest, Message: "Failed to Update Customer: bad request"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestDeleteHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().DeleteOne(gomock.Any(), 1).Return(nil).Times(1)

	app := fiber.New()

	app.Delete("/customers/:id", handler.DeleteCustomerHandler)

	req := httptest.NewRequest(http.MethodDelete, "/customers/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusOK, Message: "Delete Customer Success"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestDeleteHandler_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().DeleteOne(gomock.Any(), 1).Return(errors.New("database err")).Times(1)

	app := fiber.New()

	app.Delete("/customers/:id", handler.DeleteCustomerHandler)

	req := httptest.NewRequest(http.MethodDelete, "/customers/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusBadRequest, Message: "Failed to Delete Customer: database err"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestGetHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().FindOne(gomock.Any(), 1).Return(&Customers{
		ID:   1,
		Name: "Test",
		Age:  26,
	}, nil)

	app := fiber.New()

	app.Get("/customers/:id", handler.GetCustomerHandler)

	req := httptest.NewRequest(http.MethodGet, "/customers/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusOK, Message: "Get Customer Success", Data: &Customers{
		ID:   1,
		Name: "Test",
		Age:  26,
	}}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestGetHandler_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().FindOne(gomock.Any(), 1).Return(nil, app.ErrRecordNotFound)

	app := fiber.New()

	app.Get("/customers/:id", handler.GetCustomerHandler)

	req := httptest.NewRequest(http.MethodGet, "/customers/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusNotFound, Message: "Failed to Get Customer: record not found"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}

func TestGetHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	crud := NewMockCustomerStorage(ctrl)

	handler := NewHandler(crud)

	crud.EXPECT().FindOne(gomock.Any(), 1).Return(nil, errors.New("database error"))

	app := fiber.New()

	app.Get("/customers/:id", handler.GetCustomerHandler)

	req := httptest.NewRequest(http.MethodGet, "/customers/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedResp := Response{StatusCode: fiber.StatusBadRequest, Message: "Failed to Get Customer: database error"}
	var actualResp Response
	if err := json.Unmarshal(body, &actualResp); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, expectedResp, actualResp)
}
