package todo

//import (
//	"github.com/BETIS-Fasilkom-UI/TRYOUT-BE/internal/entities"
//	"github.com/google/uuid"
//	"gorm.io/gorm"
//)
//
//type Repository interface {
//	CreateOption(option *entities.Option) error
//	FindAllOptions() (*[]entities.Option, error)
//	FindOptionByID(id uuid.UUID) (*entities.Option, error)
//	UpdateOption(option *entities.Option) error
//	DeleteOption(id uuid.UUID) error
//}
//
//type repository struct {
//	DB *gorm.DB
//}
//
//func NewRepository(db *gorm.DB) Repository {
//	return &repository{
//		DB: db,
//	}
//}
//
//func (r *repository) CreateOption(option *entities.Option) error {
//	return r.DB.Create(option).Error
//}
//
//func (r *repository) FindAllOptions() (*[]entities.Option, error)  {
//	var options []entities.Option
//	if err := r.DB.Find(&options).Error; err != nil {
//		return nil, err
//	}
//	return &options, nil
//}
//
//func (r *repository) FindOptionByID(id uuid.UUID) (*entities.Option, error)  {
//	var option entities.Option
//	if err := r.DB.First(&option, id).Error; err != nil {
//		return nil, err
//	}
//	return &option, nil
//}
//
//func (r *repository) UpdateOption(option *entities.Option) error  {
//	return r.DB.Save(option).Error
//}
//
//func (r *repository) DeleteOption(id uuid.UUID) error {
//	var option entities.Option
//	if err := r.DB.Where("id = ?", id).First(&option).Error; err != nil {
//		return err
//	}
//	return r.DB.Delete(&option).Error
//}
