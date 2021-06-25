package alsc

// Item 
type Item struct {

    // 实际费用
    ActualFee   int64 `json:"actual_fee,omitempty"`

    // 商品名称
    GoodsName   string `json:"goods_name,omitempty"`

    // 价格
    Price   int64 `json:"price,omitempty"`

    // 优惠金额
    PromFee   int64 `json:"prom_fee,omitempty"`

    // 商品规格，需保证唯一
    Sku   string `json:"sku,omitempty"`

    // 商品码
    Spu   string `json:"spu,omitempty"`

    // 总金额
    TotalFee   int64 `json:"total_fee,omitempty"`

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 数量
    Quantity   string `json:"quantity,omitempty"`

    // 一级分类
    PrimaryClass   string `json:"primary_class,omitempty"`

    // 二级分类
    SecondaryClass   string `json:"secondary_class,omitempty"`

}
