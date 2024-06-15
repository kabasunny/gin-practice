package dto

// クライアントからサーバーへのデータ転送を助けるためのデータ転送オブジェクト（DTO）。
// この構造体は、新しいアイテムを作成するために必要なデータを含み、この構造体にはバリデーションルールも含まれており、これによりクライアントから送信されるデータが適切な形式であることが確認される。

type CreateItemInput struct {
	Name        string `json:"name" binding:"requires,min=2"`
	Price       uint   `json:"price" binding:"requires,min=1,max=999999"`
	Description string `json:"description"`
}
