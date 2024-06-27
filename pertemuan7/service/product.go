package service

import (
	"pertemuan7/model"
	"pertemuan7/repository"
)

type UserService struct {
	UserLocalRepo *repository.ProductLocalRepo
	UserPgRepo    *repository.ProductPgRepo
}

func (u *UserService) Get() ([]*model.Product, error) {
	return u.UserPgRepo.Get()
}

func (u *UserService) Create(student *model.Product) error {
	return u.UserPgRepo.Create(student)
}

func (u *UserService) Update(id uint64, student *model.ProductUpdate) error {
	return u.UserPgRepo.Update(id, student)
}

func (u *UserService) Delete(id uint64) error {
	return u.UserPgRepo.Delete(id)
}

func (u *UserService) GetByEmail(email string) (*model.User, error) {
	return u.UserPgRepo.GetByEmail(email)
}
