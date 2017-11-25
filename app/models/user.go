package models

import (
	"time"

	"github.com/zanjs/go-echo-rest/db"
)

type (
	// User is
	User struct {
		BaseModel
		Username string    `json:"username" gorm:"type:varchar(100);unique"`
		Email    string    `json:"email" gorm:"type:varchar(100);unique"`
		Password string    `json:"-"`
		Articles []Article `json:"articles" gorm:"many2many:articles;"`
	}
)

//CreateTable user
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{}, &Article{}, &Product{}, &Wareroom{}, &Record{})

	return nil
}

func GetUsers() ([]User, error) {
	var (
		users    []User
		user     User
		articles []Article
		err      error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&user).Related(&articles).Find(&users).Error; err != nil {
		tx.Rollback()
		return users, err
	}
	// tx.Model(&users).Related(&articles)

	tx.Commit()

	return users, err
}

func GetUsersAnd() ([]User, error) {
	var (
		users []User
		user  User
		// articles []Article
		err error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&users).Error; err != nil {
		tx.Rollback()
		return users, err
	}
	tx.Model(&user).Related(&user.Articles)

	tx.Commit()

	return users, err
}

func GetUserById(id uint64) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&user, id).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}

func GetUserByUsername(username string) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}

func CreateUser(m *User) error {
	var (
		err error
	)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m *User) UpdateUser(data *User) error {
	var (
		err error
	)

	m.UpdatedAt = time.Now()
	m.Username = data.Username
	m.Email = data.Email
	m.Password = data.Password

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m User) DeleteUser() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
