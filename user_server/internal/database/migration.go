package database

import (
	"gorm.io/gorm"
	"time"
)

type Ship struct {
	gorm.Model
	name     string  `gorm:"notNull;type:varchar(128)"`
	IceClass string  `gorm:"notNull"`
	Velocity float64 `gorm:"notNull"`
	Location float64 `gorm:"notNull"`
}
type Way struct {
	gorm.Model
	name        string    `gorm:"notNull;type:varchar(128)"`
	ShipID      uint      `gorm:"notNull"`
	StartTime   time.Time `gorm:"notNull"`
	EndTime     time.Time `gorm:""`
	Destination string    `gorm:"notNull"`
	Departure   string    `gorm:"notNull"`
}

type Template struct {
	gorm.Model
	Title        string `gorm:"unique;notNull;type:varchar(128)"`
	Description  string `gorm:"type:varchar(512)"`
	AuthorId     uint   `gorm:"notNull"`
	SchemaPath   string `gorm:"notNull"`
	MetaPath     string `gorm:"notNull"`
	TemplatePath string `gorm:"notNull"`

	Experiment []Experiment `gorm:"foreignKey:TemplateId"`
}

type Experiment struct {
	gorm.Model
	Title       string    `gorm:"unique;notNull;type:varchar(128)"`
	Description string    `gorm:"type:varchar(512)"`
	AuthorId    uint      `gorm:"notNull"`
	TemplateId  uint      `gorm:"notNull"`
	InputData   string    `gorm:"notNull"`
	LogFilePath string    `gorm:""`
	Status      int8      `gorm:"type:smallint;default:0"`
	StartTime   time.Time `gorm:""`
	EndTime     time.Time `gorm:""`
	SchemaPath  string    `gorm:"notNull"`
	PodPath     string    `gorm:"notNull"`

	ResultMessage []ResultMessage `gorm:"foreignKey:ExperimentId"`
}

type ResultMessage struct {
	gorm.Model
	ExperimentId uint64 `gorm:"notNull"`
	ModuleTitle  string `gorm:"notNull;type:varchar(128)"`
	Title        string `gorm:"notNull;type:varchar(128)"`
	Data         string `gorm:""`
}

func (class Template) GetID() uint {
	return class.ID
}
func (class Template) GetTitle() string {
	return class.Title
}

func (class Experiment) GetID() uint {
	return class.ID
}
func (class Experiment) GetTitle() string {
	return class.Title
}

func (class ResultMessage) GetID() uint {
	return class.ID
}
func (class ResultMessage) GetTitle() string {
	return class.Title
}
