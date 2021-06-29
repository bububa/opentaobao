package eleenterpriseordernew

// OrderExtra 
type OrderExtra struct {
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 费用
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 费用项名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 订单项目分类（参考附录）
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}