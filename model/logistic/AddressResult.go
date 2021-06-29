package logistic

// AddressResult 
type AddressResult struct {
    // 修改日期时间
    ModifyDate   string `json:"modify_date,omitempty" xml:"modify_date,omitempty"`
    // 地址库ID
    ContactId   int64 `json:"contact_id,omitempty" xml:"contact_id,omitempty"`
    // 联系人姓名
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    // 省
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 区、县
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // 详细街道地址，不需要重复填写省/市/区
    Addr   string `json:"addr,omitempty" xml:"addr,omitempty"`
    // 地区邮政编码
    ZipCode   string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
    // 电话号码,手机与电话必需有一个
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 手机号码，手机与电话必需有一个 <br/>手机号码不能超过20位
    MobilePhone   string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
    // 公司名称,
    SellerCompany   string `json:"seller_company,omitempty" xml:"seller_company,omitempty"`
    // 备注
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // 区域ID
    AreaId   int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
    // 是否默认发货地址
    SendDef   bool `json:"send_def,omitempty" xml:"send_def,omitempty"`
    // 是否默认取货地址
    GetDef   bool `json:"get_def,omitempty" xml:"get_def,omitempty"`
    // 是否默认退货地址
    CancelDef   bool `json:"cancel_def,omitempty" xml:"cancel_def,omitempty"`
}
