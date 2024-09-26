package models

import (
	"time"
)
type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password  string `json:"password"`
	Habits []Habit `json:"foreignKey:UserID"`
}
type Categorie struct{
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Habits []Habit `json:"habits" gorm:"foreignKey:CategoryID"`

}
type Goals struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Habit_id uint `json:"habitid"`
	Target int `json:"target"`
	Current int `json:"current"`
	Streak int `json:"streak"`
	Createdat time.Time `json:"date"`
	UpdatedAt time.Time `json:"updatedAt"`


}
type Habit struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Userid uint `json:"userid"`
	CategoryID uint     `json:"categoryId"`
	Name string `json:"name"`
	Createdat time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedAt"`
	Progresses []Progres `gorm:"foreignKey:Habit_id" json:"progresses"`
	Reminders []Reminder `gorm:"foreignKey:Habit_id" json:"reminders"`
}
type Progres struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Habit_id uint `json:"habitid"`
	Date time.Time `json:"date"`
	Status    string  `json:"status"`

}
type Reminder struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Habit_id uint `json:"habitid"`
	Date time.Time `json:"date"`
	Status    string  `json:"status"`
}