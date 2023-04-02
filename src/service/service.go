package service

import "src/model"

type DairyService interface {
	Save(model.Dairies) model.Dairies
	Show(model.Dairies) model.Dairies
}

type dairyService struct {
	dairy []model.Dairies
}

func New() DairyService {
	return &dairyService{}
}

func (service *dairyService) Save(dairy model.Dairies) model.Dairies {
	service.dairy = append(service.dairy, dairy)
	return dairy
}

func (service *dairyService) Show(dairy model.Dairies) model.Dairies {
	return dairy
}
