package middleclaims

import (
	"sync"
)

// ClaimsBillDto 结构体
type ClaimsBillDto struct {
	// 报案号
	ReportNo string `json:"report_no,omitempty" xml:"report_no,omitempty"`
	// 币种
	AmountCurrency string `json:"amount_currency,omitempty" xml:"amount_currency,omitempty"`
	// 收款账户
	Payee string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 服务工单ID
	ServiceWorkOrderId int64 `json:"service_work_order_id,omitempty" xml:"service_work_order_id,omitempty"`
	// 金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 预留扩展Map
	ExtensionMap *Extensionmap `json:"extension_map,omitempty" xml:"extension_map,omitempty"`
}

var poolClaimsBillDto = sync.Pool{
	New: func() any {
		return new(ClaimsBillDto)
	},
}

// GetClaimsBillDto() 从对象池中获取ClaimsBillDto
func GetClaimsBillDto() *ClaimsBillDto {
	return poolClaimsBillDto.Get().(*ClaimsBillDto)
}

// ReleaseClaimsBillDto 释放ClaimsBillDto
func ReleaseClaimsBillDto(v *ClaimsBillDto) {
	v.ReportNo = ""
	v.AmountCurrency = ""
	v.Payee = ""
	v.PayTime = ""
	v.ServiceWorkOrderId = 0
	v.Amount = 0
	v.ExtensionMap = nil
	poolClaimsBillDto.Put(v)
}
