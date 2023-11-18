package models

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Name     string
	Monsters []Monster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CampaignID;references:ID"`
}

type Monster struct {
	gorm.Model
	Name              string
	CampaignID        uint
	Type              string
	SubType           string
	Height            string
	Alignment         string
	ArmorClass        string
	HealthPoints      uint64
	Speed             uint32
	Strength          uint8
	Dexterity         uint8
	Constitution      uint8
	Intelligence      uint8
	Wisdom            uint8
	Charisma          uint8
	SavingThrows      string
	Skills            string
	Vulnerabilities   string
	Resistances       string
	DamageImmunities  string
	StateImmunities   string
	Senses            string
	PassivePerception uint8
	Languages         string
	Challenge         string
	MasteryBonus      uint8
	SpecialTraits     []SpecialTrait `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MonsterID;references:ID"`
	Actions           []Action       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MonsterID;references:ID"`
	Description       string
	Portrait          []byte
}

type SpecialTrait struct {
	gorm.Model
	Name        string
	Description string
	MonsterID   uint
}

type Action struct {
	gorm.Model
	Name        string
	Description string
	Type        string
	MonsterID   uint
}

type Height struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Alignment struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
