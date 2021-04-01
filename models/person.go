package models

type Person struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;column:Id" json:"id"`
	Deleted    bool   `gorm:"type:bit;not null;default:0;column:Deleted" json:"deleted"`
	FirstName  string `gorm:"type:nvarchar(50);not null;column:FirstName" json:"firstName"`
	MiddleName string `gorm:"type:nvarchar(50);column:MiddleName" json:"middleName"`
	LastName   string `gorm:"type:nvarchar(50);not null;column:LastName" json:"lastName"`
}

func (Person) TableName() string {
	return "People"
}

type PersonSource interface {
	GetPersons(offset, limit int) ([]Person, error)
	GetPerson(id int64) (Person, error)
	CreatePerson(person *Person) error
	UpdatePerson(person *Person) error
	DeletePerson(id int64) error
}
