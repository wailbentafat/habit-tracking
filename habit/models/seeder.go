package models
 
import (
	"gorm.io/gorm"
)

// SeedCategories seeds predefined categories into the database
func SeedCategories(db *gorm.DB) {
	categories := []Categorie{
		{Name: "Fitness"},
		{Name: "Health"},
		{Name: "Learning"},
		{Name: "Productivity"},
		{Name: "Wellness"},
	}

	for _, category := range categories {
		db.FirstOrCreate(&category, Categorie{Name: category.Name})
	}
}