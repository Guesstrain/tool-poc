package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginservice struct {
	authorizedUsername string
	authorizePassword  string
}

func NewLoginService() LoginService {
	return &loginservice{
		authorizedUsername: "pragmatic",
		authorizePassword:  "reviews",
	}
}

func (service *loginservice) Login(username string, password string) bool {
	return service.authorizedUsername == username && service.authorizePassword == password
}
