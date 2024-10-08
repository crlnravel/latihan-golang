package entities

import "gorm.io/gorm"

//type Option struct {
//	gorm.Model
//	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
//	QuestionID uuid.UUID `json:"question_id"`
//	Text       string    `json:"text"`
//}

type Todo struct {
	gorm.Model
	// TODO: ISI
}
