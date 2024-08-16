package infra

import (
	"log" //ログ出力を行うための標準ライブラリ

	"github.com/joho/godotenv" // .envファイルを読み込むための外部ライブラリ
)

func Initialize() {
	err := godotenv.Load() // .envファイルから環境変数を読み込み
	if err != nil {
		log.Fatal("Error loading .env file") // 読み込みに失敗した場合、エラーメッセージを出力し、プログラムを終了
	}
}
