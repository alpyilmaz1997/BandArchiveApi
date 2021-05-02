package bandhandlers

import (
	"github.com/alpyilmaz1997/BandArchiveApi/models"
	"github.com/alpyilmaz1997/BandArchiveApi/restapi/operations/bands"
	"github.com/alpyilmaz1997/BandArchiveApi/services/banderrors"
	"github.com/go-openapi/runtime/middleware"
)

type Bandservice interface {
	GetBands() ([]*models.Band, error)
	GetBandByID(id int64) (*models.Band, error)
	AddBand(body *models.NewBand) (*models.Band, error)
}

type Bandhandlers struct {
	service Bandservice
}

func New(bs Bandservice) *Bandhandlers {
	return &Bandhandlers{
		service: bs,
	}
}

func (bh *Bandhandlers) GetBands(params bands.GetBandsParams) middleware.Responder {
	datas, err := bh.service.GetBands()
	if err == banderrors.ErrBandNotFound {
		return bands.NewGetBandsNotFound()
	}
	if err != nil {
		return bands.NewGetBandsInternalServerError()
	}
	return bands.NewGetBandsOK().WithPayload(datas)

}
func (bh *Bandhandlers) GetBandByID(params bands.GetBandByIDParams) middleware.Responder {
	data, err := bh.service.GetBandByID(params.ID)
	if err == banderrors.ErrBandNotFound {
		return bands.NewGetBandByIDNotFound()
	}
	if err != nil {
		return bands.NewGetBandByIDInternalServerError()
	}
	return bands.NewGetBandByIDOK().WithPayload(data)
}
func (bh *Bandhandlers) AddBand(params bands.AddBandParams) middleware.Responder {
	newband, err := bh.service.AddBand(params.Body)
	if err != nil {
		return bands.NewAddBandInternalServerError()
	}
	return bands.NewAddBandOK().WithPayload(newband)
}
