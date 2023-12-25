package repositories

import (
	"errors"
	"mpc-api/config"
	"mpc-api/entities"
	"mpc-api/models"
)

func InsertUser(newUser entities.UserCore) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Role:        newUser.Role,
	}
	// simpan ke DB
	tx := config.DB.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func SelectAllUsers() ([]entities.UserCore, error) {
	var usersDataGorm []models.User
	tx := config.DB.Find(&usersDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var usersDataCore []entities.UserCore
	for _, value := range usersDataGorm {
		var userCore = entities.UserCore{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Password:    value.Password,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersDataCore = append(usersDataCore, userCore)
	}
	return usersDataCore, nil
}

func UpdateUserById(id int, userUpdate models.User) error {
	tx := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(userUpdate)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

func DeleteUserById(id int, userDelete models.User) error {
	// Menghapus pengguna dari database berdasarkan id
	tx := config.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func SelectUserById(id int) (entities.UserCore, error) {
	var userDataGorm models.User
	tx := config.DB.First(&userDataGorm, id) // select * from users where id = ?;

	if tx.Error != nil {
		return entities.UserCore{}, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	userDataCore := entities.UserCore{
		ID:          userDataGorm.ID,
		Name:        userDataGorm.Name,
		Email:       userDataGorm.Email,
		Password:    userDataGorm.Password,
		Address:     userDataGorm.Address,
		PhoneNumber: userDataGorm.PhoneNumber,
		Role:        userDataGorm.Role,
		CreatedAt:   userDataGorm.CreatedAt,
		UpdatedAt:   userDataGorm.UpdatedAt,
	}
	return userDataCore, nil
}


func ListUserProducts(id int) ([]entities.ProductCore, error) {
	var user models.User
	tx := config.DB.Preload("Products").Preload("Products.User").Find(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	// proses mapping dari struct gorm model ke struct core
	var productsDataCore []entities.ProductCore
	for _, value := range user.Products {
		var userCore = entities.UserCore{
			ID:          value.User.ID,
			Name:        value.User.Name,
			Role:        value.User.Role,
		}

		var productCore = entities.ProductCore{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			User:        userCore,
		}
		productsDataCore = append(productsDataCore, productCore)
	}
	return productsDataCore, nil
}
