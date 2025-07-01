package employees

type Employee struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}
