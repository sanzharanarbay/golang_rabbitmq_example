package models

type User struct {
	ID         int64   `json:"id" binding:"required"`
	FIO        string  `json:"fio" binding:"required"`
	Department string  `json:"department" binding:"required"`
	Age        int64   `json:"age" binding:"required"`
	Mark       float64 `json:"mark" binding:"required"`
}
