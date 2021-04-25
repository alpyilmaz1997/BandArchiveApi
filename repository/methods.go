package repository

import (
	"github.com/alpyilmaz1997/BandArchiveApi/dbmodels"
)

func (r *Repository) GetBands() (bandlist []*dbmodels.Dbband) {
	err := r.engine.Table("bands").OrderBy("id DESC").Find(&bandlist)
	if err != nil {
		panic(err)
	}
	return
}
func (r *Repository) GetBandByID(id int64) (receivedBand *dbmodels.Dbband) {
	var rec []*dbmodels.Dbband
	err := r.engine.Table("bands").Where("id = ?", id).Find(&rec)
	receivedBand = rec[0]
	if err != nil {
		panic(err)
	}
	return
}
func (r *Repository) AddBand(body *dbmodels.Dbband) {
	_, err := r.engine.Table("bands").Insert(body)
	if err != nil {
		panic(err)
	}

}
