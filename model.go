package gorm

import (
	"encoding/json"
	"log"
	"time"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExtraJson string `gorm:"type:text" json:"-"`
	Extra map[string]interface{} `gorm:"-" json:"extra"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func (m *Model) SetExtra(p map[string]interface{}) {
	m.Extra = p
	if e, err := json.Marshal(p); err != nil {
		log.Fatalf("model extra set er. p: %v, err: %v", p, err)
	} else {
		m.ExtraJson = string(e)
	}
}

func (m *Model) GetExtra() map[string]interface{} {
	if m.Extra == nil {
		m.Extra = make(map[string]interface{})
		if m.ExtraJson != "" {
			if err := json.Unmarshal([]byte(m.ExtraJson), &m.Extra); err != nil {
				log.Fatalf("model extra unmarshal fail. extraJson: %s", m.ExtraJson)
			}
		}
	}
	return m.Extra
}

func (m *Model) AddExtra(key string, value interface{}) {
	e := m.GetExtra()
	if oldValue, exist := e[key]; exist {
		log.Printf("model extra key[%s] override. oldValue: %v, newValue: %v", key, oldValue, value)
	}
	e[key] = value
	m.SetExtra(e)
}

