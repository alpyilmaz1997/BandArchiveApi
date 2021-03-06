package bandservice

import (
	"github.com/alpyilmaz1997/BandArchiveApi/dbmodels"
	"github.com/alpyilmaz1997/BandArchiveApi/models"
)

type Repository interface {
	GetBands() ([]*dbmodels.Dbband, error)
	GetBandByID(id int64) (*dbmodels.Dbband, error)
	AddBand(body *dbmodels.Dbband) (*dbmodels.Dbband, error)
}

type Bandservice struct {
	repo Repository
}

func New(repo Repository) *Bandservice {
	return &Bandservice{
		repo: repo,
	}
}

func (bs *Bandservice) GetBands() ([]*models.Band, error) {
	var allbands []*models.Band
	list, err := bs.repo.GetBands()
	if err != nil {
		return nil, err
	}
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
	return allbands, nil
}
func (bs *Bandservice) GetBandByID(id int64) (*models.Band, error) {
	received, err := bs.repo.GetBandByID(id)
	if err != nil {
		return nil, err
	}
	band := &models.Band{
		Discography: received.Discography,
		Formed:      received.Formed,
		Genre:       received.Genre,
		ID:          received.ID,
		Members:     received.Members,
		Name:        received.Name,
		Status:      received.Status,
	}
	return band, nil
}
func (bs *Bandservice) AddBand(body *models.NewBand) (*models.Band, error) {
	outgoingBand := &dbmodels.Dbband{
		Discography: body.Discography,
		Formed:      body.Formed,
		Genre:       body.Genre,
		Members:     body.Members,
		Name:        body.Name,
		Status:      body.Status,
	}
	received, err := bs.repo.AddBand(outgoingBand)
	if err != nil {
		return nil, err
	}
	newband := &models.Band{
		Discography: received.Discography,
		Formed:      received.Formed,
		Genre:       received.Genre,
		ID:          received.ID,
		Members:     received.Members,
		Name:        received.Name,
		Status:      received.Status,
	}
	return newband, nil
}
