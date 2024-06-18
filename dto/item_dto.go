package dto

// DTOとは、クライアントからサーバーへのデータ転送を助けるためのデータ転送オブジェクトのこと

// この構造体は、新しいアイテムを作成するために必要なデータを含み、この構造体にはバリデーションルールも含まれており、これによりクライアントから送信されるデータが適切な形式であることが確認される。

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

// 既存のアイテムを更新するためのDTO
type UpdateItemInput struct {
	Name *string `json:"name" binding:"omitnil,min=2"`
	// omitnilタグがあるかないかに関わらず、Go言語ではポインタ型（*string, *uint, *boolなど）のフィールドはnilを許容するが、omitnilタグがある場合とない場合では、バリデーションの挙動が異なる。
	// フィールドがnilで、omitnilタグがある場合はバリデーションチェックはスキップされる。omitnilタグがない場合はバリデーションチェックが行われる。
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
