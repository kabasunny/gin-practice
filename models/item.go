package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Item struct {
	gorm.Model         // ID unit を含む構造体となっている
	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"` // 複数の制約を記述する場合はセミコロンで区切る
	UserId      uint `gorm:"not null"`               // GORMでは、構造体名に続く ID というフィールド名がデフォルトで外部キーとして認識される
}
