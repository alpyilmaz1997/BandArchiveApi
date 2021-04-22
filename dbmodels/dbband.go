package dbmodels

type dbband struct {
	Discography []string `xorm:"text[] 'discography'"`

	Formed string `xorm:"varchar 'formed'"`

	Genre []string `xorm:"text[] 'genre'"`

	ID int64 `xorm:"bigserial pk 'id'"`

	Members []string `xorm:"text[] 'members'"`

	Name string `xorm:"varchar 'name'"`

	Status bool `xorm:"boolean 'status'"`
}
