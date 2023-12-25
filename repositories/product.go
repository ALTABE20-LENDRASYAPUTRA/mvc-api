package repositories

import (
	"errors"
	"mpc-api/config"
	"mpc-api/entities"
	"mpc-api/models"
)

func InsertProduct(newProduct entities.ProductCore) error {
	// proses mapping dari struct entities core ke model gorm
	productInputGorm := models.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		UserID:      newProduct.UserID,
	}
	// simpan ke DB
	tx := config.DB.Create(&productInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func SelectAllProducts() ([]models.Product, error) {
	var productsData []models.Product
	tx := config.DB.Preload("User").Find(&productsData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return productsData, nil
}

func UpdateProductById(id int, productUpdate models.Product) error {
	tx := config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(productUpdate)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

func DeleteProductById(id int, productDelete models.Product) error {
	// Menghapus pengguna dari database berdasarkan id
	tx := config.DB.Delete(&models.Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func SelectProductById(id int) (entities.ProductCore, error) {
	var productDataGorm models.Product
	tx := config.DB.Preload("User").First(&productDataGorm, id) // select * from products where id = ?;

	if tx.Error != nil {
		return entities.ProductCore{}, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	userCore := entities.UserCore{
		ID:          productDataGorm.User.ID,
		Name:        productDataGorm.User.Name,
		Role:        productDataGorm.User.Role,
	}

	productDataCore := entities.ProductCore{
		ID:          productDataGorm.ID,
		Name:        productDataGorm.Name,
		Description: productDataGorm.Description,
		CreatedAt:   productDataGorm.CreatedAt,
		UpdatedAt:   productDataGorm.UpdatedAt,
		User:        userCore,
	}
	return productDataCore, nil
}

