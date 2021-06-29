package openmall

// TopPostageVo 
type TopPostageVo struct {
    // 增费：支持0.00-999.99（最多包含两位小数）
    AddFee   string `json:"add_fee,omitempty" xml:"add_fee,omitempty"`
    // 增费标准。 当valuation为0时，支持1-9999范围内的整数；  当valuation为1时，支持0.1-9999.9范围内的小数，只能包含一位小数（单位为千克）；  当valuation为3时，支持0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）。
    AddStandard   string `json:"add_standard,omitempty" xml:"add_standard,omitempty"`
    // 涉及的地区，多个地区用逗号分隔。地区码可以用taobao.areas.get接口获取，或者参考：http://www.mca.gov.cn/article/sj/xzqh/2020/2020/202003301019.html  当值为1时，表示全国所有地区。
    PostArea   string `json:"post_area,omitempty" xml:"post_area,omitempty"`
    // 运费方式，可选值：平邮 (post)、快递公司(express)、EMS (ems)、货到付款(cod)、物流宝保障速递(wlb)、家装物流(furniture)。
    PostType   string `json:"post_type,omitempty" xml:"post_type,omitempty"`
    // 首费，范围0.00-999.99（最多包含两位小数）。
    StartFee   string `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
    // 首费标准。  当valuation(记价方式)为1时，支持0.1-9999.9范围内的小数只能包含一位小数（单位为千克）；  当valuation(记价方式)为3时，支持0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）。
    StartStandard   string `json:"start_standard,omitempty" xml:"start_standard,omitempty"`
    // 运费计算方式，可选值：0（件数）、1（重量）、3（体积）。
    Valuation   int64 `json:"valuation,omitempty" xml:"valuation,omitempty"`
}
