package bandservice

import (
	"github.com/alpyilmaz1997/BandArchiveApi/dbmodels"
	"github.com/alpyilmaz1997/BandArchiveApi/models"
)

type Repository interface {
	GetBands() []*dbmodels.Dbband
	GetBandByID(id int64) *dbmodels.Dbband
	AddBand(body *dbmodels.Dbband)
}

type Bandservice struct {
	repo Repository
}

func New(repo Repository) *Bandservice {
	return &Bandservice{
		repo: repo,
	}
}

func (bs *Bandservice) GetBands() []*models.Band {
	var allbands []*models.Band
	list := bs.repo.GetBands()
	for _, band := range list {
		allbands = append(allbands, &models.Band{
			Discography: band.Discography,
			Formed:      band.Formed,
			Genre:       band.Genre,
			ID:          band.ID,
			Members:     band.Members,
			Name:        band.Name,
			Status:      band.Status,
		})
	}
	return allbands
}
func (bs *Bandservice) GetBandByID(id int64) *models.Band {
	received := bs.repo.GetBandByID(id)
	band := &models.Band{
		Discography: received.Discography,
		Formed:      received.Formed,
		Genre:       received.Genre,
		ID:          received.ID,
		Members:     received.Members,
		Name:        received.Name,
		Status:      received.Status,
	}
	return band
}
func (bs *Bandservice) AddBand(body *models.NewBand) {
	outgoingBand := &dbmodels.Dbband{
		Discography: body.Discography,
		Formed:      body.Formed,
		Genre:       body.Genre,
		Members:     body.Members,
		Name:        body.Name,
		Status:      body.Status,
	}
	bs.repo.AddBand(outgoingBand)
}
