package utils

// Кастомная ошибка для приложения

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
}
