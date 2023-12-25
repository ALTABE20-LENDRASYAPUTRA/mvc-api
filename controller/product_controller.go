package controller

import (
	"mpc-api/entities"
	"mpc-api/models"
	"mpc-api/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	newProduct := entities.ProductCore{}
	errBind := c.Bind(&newProduct) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	// simpan ke DB
	errInsert := repositories.InsertProduct(newProduct)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "create data success",
	})
}

func GetAllProductController(c echo.Context) error {
	results, errSelect := repositories.SelectAllProducts()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error read data. " + errSelect.Error(),
		})
	}

	// proses mapping dari core ke response
	var productsResult []entities.ProductResponse
	for _, value := range results {
		productResponse := entities.ProductResponse{
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

		productsResult = append(productsResult, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all data",
		"data":    productsResult,
	})
}

func UpdateProductByIdController(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	var ProductData = models.Product{}
	errBind := c.Bind(&ProductData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	err := repositories.UpdateProductById(idParam, ProductData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update product",
	})
}


func DeleteProductByIdController(c echo.Context) error {
	// mendapatkan id dari parameter
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)

	// Jika terjadi kesalahan dalam konversi id, kembalikan pesan kesalahan
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	var productData = models.Product{}
	errBind := c.Bind(&productData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	err := repositories.DeleteProductById(idParam, productData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete data",
	})
}

func GetProductByIdController(c echo.Context) error {
	// mendapatkan id dari parameter
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)

	// Jika terjadi kesalahan dalam konversi id, kembalikan pesan kesalahan
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	product, errSelect := repositories.SelectProductById(idParam)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data. " + errSelect.Error(),
		})
	}

	// proses mapping dari core ke response
	productResponse := entities.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		User: entities.UserProResponse{
			ID: product.User.ID,
			Name:        product.User.Name,
			Role:        product.User.Role,
		},
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get data",
		"data":    productResponse,
	})
}