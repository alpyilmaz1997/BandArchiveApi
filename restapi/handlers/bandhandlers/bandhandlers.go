package bandhandlers

import (
	"github.com/alpyilmaz1997/BandArchiveApi/models"
	"github.com/alpyilmaz1997/BandArchiveApi/restapi/operations/bands"
	"github.com/go-openapi/runtime/middleware"
)

type Bandservice interface {
	GetBands() []*models.Band
	GetBandByID(id int64) *models.Band
	AddBand(body *models.NewBand)
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
	datas := bh.service.GetBands()
	return bands.NewGetBandsOK().WithPayload(datas)

}
func (bh *Bandhandlers) GetBandByID(params bands.GetBandByIDParams) middleware.Responder {
	data := bh.service.GetBandByID(params.ID)
	return bands.NewGetBandByIDOK().WithPayload(data)
}
func (bh *Bandhandlers) AddBand(params bands.AddBandParams) middleware.Responder {
	bh.service.AddBand(params.Body)
	return bands.NewAddBandOK()
}
