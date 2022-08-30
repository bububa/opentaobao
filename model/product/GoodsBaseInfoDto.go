package product

// GoodsBaseInfoDto 结构体
type GoodsBaseInfoDto struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 库存
	Storage int64 `json:"storage,omitempty" xml:"storage,omitempty"`
}
