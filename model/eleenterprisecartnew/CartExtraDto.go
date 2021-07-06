package eleenterprisecartnew

// CartExtraDto 结构体
type CartExtraDto struct {
	// 费用
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 费用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 费用id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单项目分类
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
