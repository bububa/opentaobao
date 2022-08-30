package ascp

// ContactInfo 结构体
type ContactInfo struct {
	// 联系人分组，string（50），mr=默认组，th=退货组，kf=客服组，rk=入库组，ck=出库组，kn=库内组
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 联系人名称，string（20）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人手机，string（20)
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 固定电话 ，string（20)
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 省份，string（15)
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市，string（15）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 地区，string（15）
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 乡镇，string（15）
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址，string（50）
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}
