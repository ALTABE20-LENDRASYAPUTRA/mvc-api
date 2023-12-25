package controller

import (
	"fmt"
	"mpc-api/entities"
	"mpc-api/models"
	"mpc-api/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	newUser := entities.UserCore{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	// simpan ke DB
	errInsert := repositories.InsertUser(newUser)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "create data success",
	})
}

func GetAllUserController(c echo.Context) error {
	results, errSelect := repositories.SelectAllUsers()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data. " + errSelect.Error(),
		})
	}
	fmt.Println("users:", results)
	// proses mapping dari core ke response
	var usersResult []entities.UserResponse
	for _, value := range results {
		usersResult = append(usersResult, entities.UserResponse{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all data",
		"data":    usersResult,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}
	var userData = models.User{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}
	err := repositories.UpdateUserById(idParam, userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update",
	})
}

func DeleteUserByIdController(c echo.Context) error {
	// mendapatkan id dari parameter
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)

	// Jika terjadi kesalahan dalam konversi id, kembalikan pesan kesalahan
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	var userData = models.User{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	err := repositories.DeleteUserById(idParam, userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete data",
	})
}


func GetUserByIdController(c echo.Context) error {
	// mendapatkan id dari parameter
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)

	// Jika terjadi kesalahan dalam konversi id, kembalikan pesan kesalahan
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	user, errSelect := repositories.SelectUserById(idParam)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data. " + errSelect.Error(),
		})
	}

	// proses mapping dari core ke response
	userResult := entities.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get data",
		"data":    userResult,
	})
}


func GetUserProductController(c echo.Context) error {
	// mendapatkan id dari parameter
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)

	// Jika terjadi kesalahan dalam konversi id, kembalikan pesan kesalahan
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	// Memanggil fungsi repositori untuk mendapatkan produk berdasarkan ID pengguna
	products, err := repositories.ListUserProducts(idParam)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data. " + err.Error(),
		})
	}

	// proses mapping dari core ke response
	var productsResult []entities.ProductResponse
	for _, value := range products {
		userProductResponse := entities.ProductResponse{
			Name:        value.Name,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			User: entities.UserProResponse{
				ID: value.User.ID,
				Name:        value.User.Name,
				Role:        value.User.Role,
			},
		}

		productsResult = append(productsResult, userProductResponse)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all data",
		"data":    productsResult,
	})
}
