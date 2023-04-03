package authentication

import (
	"Curhatku/backend/models"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Register(regData RegisterData) (string, error)
	Login(req DataRequest) (string, error)
	ChangePassword(pass Password) error
	AuthenticateUser(cookie string) (models.UserTab, error)
}

type repository struct {
	db *gorm.DB
}

const SecretKey = "secret"

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(regData RegisterData) (string, error) {
	var acc models.AppCustCodeTab
	t := time.Now()

	res2 := r.db.First(&acc)
	if res2.Error != nil {
		return "", res2.Error
	}

	//Bikin UserID
	UserID := "800"
	r.db.Model(&acc).Where("id = ?", acc.ID).Update("value", acc.Value+1)
	userIDSeq := "0000000000" + strconv.Itoa(acc.Value)
	userIDSeq = userIDSeq[len(userIDSeq)-10:]
	newUserID := UserID + regData.Birthdate.Format("20060131") + t.Format("200601") + userIDSeq

	password, _ := bcrypt.GenerateFromPassword([]byte(regData.Password), 14)

	user := models.UserTab{
		UserID:    newUserID,
		Username:  regData.Username,
		Email:     regData.Email,
		Password:  password,
		Gender:    regData.Gender,
		BirthDate: regData.Birthdate,
	}

	r.db.Create(&user)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    regData.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *repository) Login(req DataRequest) (string, error) {
	var user models.UserTab
	res := r.db.Table("user_Tab").Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return "", errors.New("email tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return "", errors.New("password salah")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	// var acc models.AppCustCodeTab
	// res2 = r.db.First(&acc)
	// if res2.Error != nil {
	// 	return nil, res.Error
	// }
	// AppCustCode := "006"
	// r.db.Model(&acc).Where("id = ?", acc.ID).Update("value", acc.Value+1)
	// userIDSeq := "0000000000" + strconv.Itoa(acc.Value)
	// userIDSeq = userIDSeq[len(userIDSeq)-10:]
	// newCustCode := AppCustCode + com[0].CompanyCode + t.Format("200601") + userIDSeq

	return token, nil
}

func (r *repository) ChangePassword(pass Password) error {
	var user models.UserTab
	res := r.db.Table("user_Tab").Where("email = ?", pass.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return res.Error
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(pass.OldPassword)); err != nil {
		return errors.New("Password tidak sama")
	} else {
		passwordBaru, _ := bcrypt.GenerateFromPassword([]byte(pass.NewPassword), 14)
		res := r.db.Table("user_Tab").Where("email = ?", pass.Email).Update("password", passwordBaru)
		if res.Error != nil {
			log.Println("Update Data error : ", res.Error)
			return res.Error
		}
	}

	return nil
}

func (r *repository) AuthenticateUser(cookie string) (models.UserTab, error) {
	var user models.UserTab

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return user, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	r.db.Table("user_Tab").Where("email = ?", claims.Issuer).First(&user)

	return user, nil
}
