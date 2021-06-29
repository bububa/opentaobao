package lstbm

// LstTopOpenStoreDTO 
type LstTopOpenStoreDTO struct {
    // 店主的手机号
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 门店名称，最大长度512字符
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 企业自定义的门店二级类别，最大长度512字符
    SecondCategory   string `json:"second_category,omitempty" xml:"second_category,omitempty"`
    // 企业自定义的门店一级类别，最大长度512字符
    PrimaryCategory   string `json:"primary_category,omitempty" xml:"primary_category,omitempty"`
    // 营业执照上的经营者名称，最大长度256字符
    LicenseOwner   string `json:"license_owner,omitempty" xml:"license_owner,omitempty"`
    // 详细地址，最大长度512字符
    Addr   string `json:"addr,omitempty" xml:"addr,omitempty"`
    // 市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 营业执照编号，最大长度256字符
    LicenseNo   string `json:"license_no,omitempty" xml:"license_no,omitempty"`
    // 店主姓名，最大长度256字符
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 省
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 区县内的街道
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    // 企业自定义门店id，最大长度128字符
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
