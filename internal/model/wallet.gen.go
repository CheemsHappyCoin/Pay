// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameWallet = "wallet"

// Wallet mapped from table <wallet>
type Wallet struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MID       int64          `gorm:"column:m_id;not null" json:"m_id"`
	Ac        string         `gorm:"column:ac;not null;comment:钱包地址" json:"ac"`            // 钱包地址
	PriKey    string         `gorm:"column:pri_key;not null;comment:私钥" json:"pri_key"`    // 私钥
	Mnemonic  string         `gorm:"column:mnemonic;not null;comment:助词器" json:"mnemonic"` // 助词器
	Path      string         `gorm:"column:path;not null;comment:钱包路径" json:"path"`        // 钱包路径
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`               // 备注
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Wallet's table name
func (*Wallet) TableName() string {
	return TableNameWallet
}
