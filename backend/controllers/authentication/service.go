package authentication

import (
	"Curhatku/backend/models"
	"log"
	"net/http"
)

type Service interface {
	Register(regData RegisterData) (int, string, error)
	Login(req DataRequest) (int, string, error)
	ChangePassword(pass Password) (int, error)
	AuthenticateUser(cookie string) (int, models.UserTab, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Register(regData RegisterData) (int, string, error) {
	token, err := s.repo.Register(regData)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, token, err
	}
	return http.StatusOK, token, nil
}

func (s *service) Login(req DataRequest) (int, string, error) {
	token, err := s.repo.Login(req)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, token, err
	}
	return http.StatusOK, token, nil
}

func (s *service) ChangePassword(pass Password) (int, error) {
	err := s.repo.ChangePassword(pass)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *service) AuthenticateUser(cookie string) (int, models.UserTab, error) {
	user, err := s.repo.AuthenticateUser(cookie)
	if err != nil {
		log.Println("Unauthorized : ", err)
		return http.StatusUnauthorized, user, err
	}
	return http.StatusOK, user, nil
}
