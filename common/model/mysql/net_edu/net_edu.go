package model

// NetEdu [...]
type NetEdu struct {
	ID     int64  `gorm:"primaryKey;column:id" json:"-"`
	Number int64  `gorm:"column:number" json:"number"`
	Body   string `gorm:"column:body" json:"body"`
}

// TableName get sql table name.获取数据库表名
func (m *NetEdu) TableName() string {
	return "net_edu"
}

// NetEduColumns get sql column name.获取数据库列名
var NetEduColumns = struct {
	ID     string
	Number string
	Body   string
}{
	ID:     "id",
	Number: "number",
	Body:   "body",
}
