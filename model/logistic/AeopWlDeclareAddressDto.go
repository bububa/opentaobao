package logistic

// AeopWlDeclareAddressDto 结构体
type AeopWlDeclareAddressDto struct {
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 类型
	MemberType string `json:"member_type,omitempty" xml:"member_type,omitempty"`
	// 旺旺
	TrademanageId string `json:"trademanage_id,omitempty" xml:"trademanage_id,omitempty"`
	// 街道
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	County string `json:"county,omitempty" xml:"county,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 详细地址
	StreetAddress string `json:"street_address,omitempty" xml:"street_address,omitempty"`
	// 电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 仓Code
	FromWarehouseCode string `json:"from_warehouse_code,omitempty" xml:"from_warehouse_code,omitempty"`
	// 卖家后台地址id,用来获取卖家详细地址信息，传入值为Long型；传入addressId后，其余字段值无效。
	AddressId int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
}
