package repository

import (
	"github.com/alpyilmaz1997/BandArchiveApi/dbmodels"
	"github.com/alpyilmaz1997/BandArchiveApi/repository/repoerrs"
)

func (r *Repository) GetBands() ([]*dbmodels.Dbband, error) {
	var bandlist []*dbmodels.Dbband
	err := r.engine.Table("bands").OrderBy("id DESC").Find(&bandlist)
	if err != nil {
		return nil, err
	}
	if len(bandlist) == 0 {
		return nil, repoerrs.ErrNotFound
	}
	return bandlist, nil
}

func (r *Repository) GetBandByID(id int64) (*dbmodels.Dbband, error) {
	var receivedBand []*dbmodels.Dbband
	err := r.engine.Table("bands").Where("id = ?", id).Find(&receivedBand)
	if err != nil {
		return nil, err
	}
	if len(receivedBand) == 0 {
		return nil, repoerrs.ErrNotFound
	}
	return receivedBand[0], nil
}

func (r *Repository) AddBand(body *dbmodels.Dbband) (*dbmodels.Dbband, error) {
	_, err := r.engine.Table("bands").Insert(body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
