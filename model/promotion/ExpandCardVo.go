package promotion

// ExpandCardVo 
type ExpandCardVo struct {
    // 店铺信息
    ShopInfoVo   *ShopInfoVo `json:"shop_info_vo,omitempty" xml:"shop_info_vo,omitempty"`
    // 品牌色
    BrandColor   string `json:"brand_color,omitempty" xml:"brand_color,omitempty"`
    // 目标跳转链接
    TargetUrl   string `json:"target_url,omitempty" xml:"target_url,omitempty"`
    // 聚合的卡当中第一张卡的开卡时间
    CardCreateDate   string `json:"card_create_date,omitempty" xml:"card_create_date,omitempty"`
    // 最近要过期的余额有效期描述
    CardValidDesc   string `json:"card_valid_desc,omitempty" xml:"card_valid_desc,omitempty"`
    // 聚合的卡当中最近要过期的卡的有效期
    CardValidDate   string `json:"card_valid_date,omitempty" xml:"card_valid_date,omitempty"`
    // 卡膨胀金余额，买家在该店铺下该卡类型的所有卡实例的总膨胀金余额
    CardRemainExpandMoney   string `json:"card_remain_expand_money,omitempty" xml:"card_remain_expand_money,omitempty"`
    // 卡本金余额，买家在该店铺下该卡类型的所有卡实例的总本金余额
    CardRemainBasicMoney   string `json:"card_remain_basic_money,omitempty" xml:"card_remain_basic_money,omitempty"`
    // 卡余额，买家在该店铺下该卡类型的所有卡实例的总余额
    CardRemainMoney   string `json:"card_remain_money,omitempty" xml:"card_remain_money,omitempty"`
    // 购物金店铺icon
    CardIconUrl   string `json:"card_icon_url,omitempty" xml:"card_icon_url,omitempty"`
    // 卡名称
    CardName   string `json:"card_name,omitempty" xml:"card_name,omitempty"`
    // 卡使用范围
    CardUsedScope   string `json:"card_used_scope,omitempty" xml:"card_used_scope,omitempty"`
}
