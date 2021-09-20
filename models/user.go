package models

import (
  "github.com/jinzhu/gorm"
)

type RegisterRequest struct {
  User User `json:"user"`
}

type PinRequest struct {
  RaspPin string `form:"pin" json:"pin" binding:"required"`
  UserID string `form:"user_id" json:"user_id" binding:"required"`
}

type User struct {
  gorm.Model
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Password string `json:"password"`
  Email string `json:"email"`
  PhoneNumber string `json:"phone_number"`
  Pin string
}
