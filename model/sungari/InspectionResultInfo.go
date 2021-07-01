package sungari

// InspectionResultInfo 结构体
type InspectionResultInfo struct {
	// 卖家nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 发货地址
	SendAddress string `json:"send_address,omitempty" xml:"send_address,omitempty"`
	// 注册地址
	RegisterAddress string `json:"register_address,omitempty" xml:"register_address,omitempty"`
	// 营业执照编号
	LicenceNo string `json:"licence_no,omitempty" xml:"licence_no,omitempty"`
	// 卖家电话
	SellerTel string `json:"seller_tel,omitempty" xml:"seller_tel,omitempty"`
	// 认证名称
	CertificationName string `json:"certification_name,omitempty" xml:"certification_name,omitempty"`
	// 处置结果
	DisposeResult string `json:"dispose_result,omitempty" xml:"dispose_result,omitempty"`
}
