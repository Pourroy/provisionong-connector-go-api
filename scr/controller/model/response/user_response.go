package response

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/request"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	Age   int8      `json:"age"`
}

func NewUserResponse(userRequest request.UserRequest) *UserResponse {
	return &UserResponse{
		ID:    uuid.New(),
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
}
