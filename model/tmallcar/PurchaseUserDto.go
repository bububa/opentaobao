package tmallcar

// PurchaseUserDto 结构体
type PurchaseUserDto struct {
	// 购车人姓名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 购车人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 购车人证号号
	IdentityNo string `json:"identity_no,omitempty" xml:"identity_no,omitempty"`
	// 证件类型编码
	IdentityType string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// 证件类型描述
	IdentityTypeDesc string `json:"identity_type_desc,omitempty" xml:"identity_type_desc,omitempty"`
}
