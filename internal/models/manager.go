package models

import (
	"github.com/google/uuid"
)

type Manager struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Number    uint      `gorm:"unique"`
	Name      string
	Functions []ManagerFunction `gorm:"foreignKey:ManagerID"`
	//Object    []Object
}

type Object struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Number    uint      `gorm:"unique"`
	Type      string
	Functions []ManagerFunction `gorm:"foreignKey:ObjectID"`
	//Manager   []Manager
}

type ManagerFunction struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Function       string
	FunctionType   string
	FunctionNumber uint
	ManagerID      uuid.UUID
	ObjectID       uuid.UUID
	Manager        Manager
	Object         Object
}

type Relations struct {
	ManagerID     uuid.UUID `gorm:"primaryKey"`
	ObjectID      uuid.UUID `gorm:"primaryKey"`
	ManagerFuncID uuid.UUID `gorm:"primaryKey"`
}
