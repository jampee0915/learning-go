package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@tcp(mysql:3306)/xxxxx")
	if err != nil {
		panic(err)
	}
}

// Corporates構造体を定義
type Corporates struct {
	CorporateId   int64  `json:"corporateId" xorm:"'corporate_id'"`
	CorporateName string `json:"corporateName" xorm:"corporate_name"`
}

func NewCorporates(corporateId int64, corporateName string) Corporates {
	return Corporates{
		CorporateId:   corporateId,
		CorporateName: corporateName,
	}
}

type CorporatesRepository struct {
}

func NewCorporatesRepository() CorporatesRepository {
	return CorporatesRepository{}
}

func (m CorporatesRepository) GetByCorporateId(corporateId int64) *Corporates {
	var corporate = Corporates{CorporateId: corporateId}
	has, _ := engine.Get(&corporate)
	if has {
		return &corporate
	}
	return nil
}
