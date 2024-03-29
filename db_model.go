package gormDBPlus

import (
	"gorm.io/gorm"
)

type Connector struct {
	RawConnection *gorm.DB
}

type ConnectionConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	SSL      bool
}

func NewConnectorWithDB(db *gorm.DB) *Connector {
	return &Connector{RawConnection: db}
}

func (c *Connector) FetchOne(dest interface{}, tx *gorm.DB, with ...string) error {
	query := getQuery(c, tx)
	query = c.RawConnection.Model(dest)
	for _, v := range with {
		query = query.Preload(v)
	}
	return query.First(dest, dest).Error
}

func getQuery(c *Connector, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return c.RawConnection
}

func (c *Connector) FetchOneWithID(dest interface{}, id int, tx *gorm.DB, with ...string) error {
	query := getQuery(c, tx)
	query = c.RawConnection
	for _, v := range with {
		query = query.Preload(v)
	}
	return query.First(dest, id).Error
}

func (c *Connector) FetchAll(dest interface{}, condition interface{}, tx *gorm.DB, with ...string) error {
	query := getQuery(c, tx)
	query = c.RawConnection
	for _, v := range with {
		query = query.Preload(v)
	}
	return query.Where(condition).Find(dest).Error
}

func (c *Connector) Save(p interface{}, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Save(p).Error
}

func (c *Connector) AppendChild(p interface{}, typeName string, child interface{}, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Model(p).Association(typeName).Append(child)
}

func (c *Connector) ReplaceChild(p interface{}, typeName string, child interface{}, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Model(p).Association(typeName).Replace(child)
}

func (c *Connector) ClearChild(p interface{}, typeName string, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Model(p).Association(typeName).Clear()
}

func (c *Connector) CountChild(p interface{}, typeName string, tx *gorm.DB) int64 {
	query := getQuery(c, tx)
	return query.Model(p).Association(typeName).Count()
}

func (c *Connector) Update(p interface{}, update interface{}, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Model(p).Updates(update).Error
}

func (c *Connector) Remove(p interface{}, tx *gorm.DB) error {
	query := getQuery(c, tx)
	return query.Model(p).Delete(p).Error
}

func (c *Connector) IsExist(p interface{}, tx *gorm.DB) (bool, error) {
	query := getQuery(c, tx)
	var count int64
	if err := query.Model(p).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

func (c *Connector) NewTransaction() *gorm.DB {
	return c.RawConnection.Begin()
}
