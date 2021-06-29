package wdk

// PaiyangStatDataResult 
type PaiyangStatDataResult struct {
    // 统计时间
    StatDate   string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
    // 活动id
    ActivityId   string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 活动名称
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    // 优惠券id
    TemplateCode   string `json:"template_code,omitempty" xml:"template_code,omitempty"`
    // 优惠券名称
    CouponName   string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 商家名称
    MerchantName   string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
    // 子公司编码
    SubsidiaryCode   string `json:"subsidiary_code,omitempty" xml:"subsidiary_code,omitempty"`
    // 子公司名称
    SubsidiaryName   string `json:"subsidiary_name,omitempty" xml:"subsidiary_name,omitempty"`
    // 经营店编码
    ShopCode   string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
    // 经营点名称
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    // 69码列表
    BarcodeList   string `json:"barcode_list,omitempty" xml:"barcode_list,omitempty"`
    // 当天使用优惠券张数
    PosUseCpnCnt1d   string `json:"pos_use_cpn_cnt1d,omitempty" xml:"pos_use_cpn_cnt1d,omitempty"`
    // 按天分区字段
    Ds   string `json:"ds,omitempty" xml:"ds,omitempty"`
}
