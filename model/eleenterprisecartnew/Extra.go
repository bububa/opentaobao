package eleenterprisecartnew

// Extra 
type Extra struct {
    // 优惠券id
    IdStr   string `json:"id_str,omitempty" xml:"id_str,omitempty"`
    // 原价
    Total   string `json:"total,omitempty" xml:"total,omitempty"`
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 费用
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 商品id
    FoodId   string `json:"food_id,omitempty" xml:"food_id,omitempty"`
    // 费用名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 减
    Reduced   int64 `json:"reduced,omitempty" xml:"reduced,omitempty"`
    // 费用项Id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 订单项目分类
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
