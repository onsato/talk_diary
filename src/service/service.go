package service

import "main/model"

type DairyService interface {
	Save(model.Diary) model.Diary
	Show(model.Diary) model.Diary
}

type dairyService struct {
	dairy []model.Diary
}

func New() DairyService {
	return &dairyService{}
}

func (service *dairyService) Save(dairy model.Diary) model.Diary {
	service.dairy = append(service.dairy, dairy)
	return dairy
}

func (service *dairyService) Show(dairy model.Diary) model.Diary {
	return dairy
}
