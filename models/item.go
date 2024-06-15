package models

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Item struct {
	ID          uint
	Name        string
	Price       uint
	Description string
	SoldOut     bool
}
