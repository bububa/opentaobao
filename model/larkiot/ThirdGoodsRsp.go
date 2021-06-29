package larkiot

// ThirdGoodsRsp 
type ThirdGoodsRsp struct {
    // 卖品编码
    GoodsCode   string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
    // 卖品名称
    GoodsName   string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
    // 套餐标识
    PackageFlag   string `json:"package_flag,omitempty" xml:"package_flag,omitempty"`
    // 标准价
    StandardPrice   string `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
    // 结算价
    SettlePrice   string `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
    // 描述
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 唯一键
    GoodsKey   string `json:"goods_key,omitempty" xml:"goods_key,omitempty"`
    // 会员标识
    MemberFlag   string `json:"member_flag,omitempty" xml:"member_flag,omitempty"`
    // 卖品url
    GoodsPicUrl   string `json:"goods_pic_url,omitempty" xml:"goods_pic_url,omitempty"`
}
