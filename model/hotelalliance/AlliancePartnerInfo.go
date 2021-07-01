package hotelalliance

// AlliancePartnerInfo 结构体
type AlliancePartnerInfo struct {
	// 合作商联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 结算账户名
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 合作商联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 合作商ID
	PartnerId int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 结算类型（目前只有0，企业支付宝）
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 结算账号
	AccountNum string `json:"account_num,omitempty" xml:"account_num,omitempty"`
	// 是否生效（0: 失效 1:生效）
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 公司联系邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 公司地址
	CompanyAddress string `json:"company_address,omitempty" xml:"company_address,omitempty"`
	// 合作商签约主账号
	MainAccount string `json:"main_account,omitempty" xml:"main_account,omitempty"`
}
