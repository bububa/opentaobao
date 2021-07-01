package alsc

// StoreKeeperDto 结构体
type StoreKeeperDto struct {
	// 门店联系人
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
}
