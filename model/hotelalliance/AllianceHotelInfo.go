package hotelalliance

import (
	"sync"
)

// AllianceHotelInfo 结构体
type AllianceHotelInfo struct {
	// 邀约人
	Inviter string `json:"inviter,omitempty" xml:"inviter,omitempty"`
	// 酒店联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 合同签署开始时间
	ContractStart string `json:"contract_start,omitempty" xml:"contract_start,omitempty"`
	// 底价比率值
	Fee string `json:"fee,omitempty" xml:"fee,omitempty"`
	// 结算账号户名
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 开户行名称
	AccountOpeningBank string `json:"account_opening_bank,omitempty" xml:"account_opening_bank,omitempty"`
	// 酒店联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 银行卡账号
	BankAccount string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
	// 合同结束时间
	ContractEnd string `json:"contract_end,omitempty" xml:"contract_end,omitempty"`
	// 酒店联系邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 银行省份
	BankProvince string `json:"bank_province,omitempty" xml:"bank_province,omitempty"`
	// 银行城市
	BankCity string `json:"bank_city,omitempty" xml:"bank_city,omitempty"`
	// 酒店签约主账号
	MainAccount string `json:"main_account,omitempty" xml:"main_account,omitempty"`
	// 开户银行
	AccountBank string `json:"account_bank,omitempty" xml:"account_bank,omitempty"`
	// 新费率
	NewFee string `json:"new_fee,omitempty" xml:"new_fee,omitempty"`
	// 营业执照公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 支付宝帐号ID
	AlipayAccountId string `json:"alipay_account_id,omitempty" xml:"alipay_account_id,omitempty"`
	// 卖家酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 合作商Id
	PartnerId int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 标准酒店Id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 费率类型
	FeeType int64 `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 结算类型0:个人支付宝 1:企业支付宝 3:对公银行卡
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 是否生效 0:失效 1:生效
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 单体联盟新生成的卖家酒店Id
	Ahid int64 `json:"ahid,omitempty" xml:"ahid,omitempty"`
	// 合同号
	TplId int64 `json:"tpl_id,omitempty" xml:"tpl_id,omitempty"`
	// 卖家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 是否菲住酒店有效
	IsNewValid int64 `json:"is_new_valid,omitempty" xml:"is_new_valid,omitempty"`
	// 新费率类型
	NewFeeType int64 `json:"new_fee_type,omitempty" xml:"new_fee_type,omitempty"`
	// 是否迁移成功
	IsTransfer int64 `json:"is_transfer,omitempty" xml:"is_transfer,omitempty"`
	// 主合同id
	MainTplId int64 `json:"main_tpl_id,omitempty" xml:"main_tpl_id,omitempty"`
}

var poolAllianceHotelInfo = sync.Pool{
	New: func() any {
		return new(AllianceHotelInfo)
	},
}

// GetAllianceHotelInfo() 从对象池中获取AllianceHotelInfo
func GetAllianceHotelInfo() *AllianceHotelInfo {
	return poolAllianceHotelInfo.Get().(*AllianceHotelInfo)
}

// ReleaseAllianceHotelInfo 释放AllianceHotelInfo
func ReleaseAllianceHotelInfo(v *AllianceHotelInfo) {
	v.Inviter = ""
	v.ContactPhone = ""
	v.ContractStart = ""
	v.Fee = ""
	v.AccountName = ""
	v.HotelName = ""
	v.AccountOpeningBank = ""
	v.ContactName = ""
	v.AlipayAccount = ""
	v.BankAccount = ""
	v.ContractEnd = ""
	v.ContactEmail = ""
	v.BankProvince = ""
	v.BankCity = ""
	v.MainAccount = ""
	v.AccountBank = ""
	v.NewFee = ""
	v.CompanyName = ""
	v.AlipayAccountId = ""
	v.Hid = 0
	v.PartnerId = 0
	v.Shid = 0
	v.FeeType = 0
	v.PayType = 0
	v.IsValid = 0
	v.Ahid = 0
	v.TplId = 0
	v.SellerId = 0
	v.IsNewValid = 0
	v.NewFeeType = 0
	v.IsTransfer = 0
	v.MainTplId = 0
	poolAllianceHotelInfo.Put(v)
}
