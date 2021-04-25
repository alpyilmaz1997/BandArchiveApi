package dbmodels

import "github.com/lib/pq"

type Dbband struct {
	Discography pq.StringArray `xorm:"text[] 'discography'"`

	Formed string `xorm:"varchar 'formed'"`

	Genre pq.StringArray `xorm:"text[] 'genre'"`

	ID int64 `xorm:"bigserial pk 'id'"`

	Members pq.StringArray `xorm:"text[] 'members'"`

	Name string `xorm:"varchar 'name'"`

	Status bool `xorm:"boolean 'status'"`
}
