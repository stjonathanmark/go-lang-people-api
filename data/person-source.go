package data

import (
	"gorm.io/gorm"
	"stjonathanmark.com/people/models"
)

func NewPersonSource(db *gorm.DB) *PersonSource {
	return &PersonSource{db, 0}
}

type PersonSource struct {
	*gorm.DB
	ItemCount int64
}

func (s *PersonSource) GetPersons(offset, limit int) (persons []models.Person, err error) {
	s.ItemCount = 0
	s.Model(&models.Person{}).Where("Deleted = ?", false).Count(&s.ItemCount)
	if limit == 0 {
		limit = int(s.ItemCount)
	}

	result := s.Where("Deleted = ?", false).Offset(offset).Limit(limit).Find(&persons)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (s *PersonSource) GetPerson(id int64) (person models.Person, err error) {
	result := s.Where("Deleted = ?", 0).First(&person, id)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (s *PersonSource) CreatePerson(person *models.Person) (err error) {
	result := s.Create(&person)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (s *PersonSource) UpdatePerson(person *models.Person) (err error) {
	per := models.Person{Id: person.Id}
	result := s.Model(&per).Omit("Id").Updates(person)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (s *PersonSource) DeletePerson(id int64) (err error) {
	person, err := s.GetPerson(id)

	if err != nil {
		return
	}

	person.Deleted = true

	result := s.Save(&person)

	if result.Error != nil {
		err = result.Error
	}

	return
}
