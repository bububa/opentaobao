package middleclaims

import (
	"sync"
)

// ClaimsResultDto 结构体
type ClaimsResultDto struct {
	// 报案号
	ReportNo string `json:"report_no,omitempty" xml:"report_no,omitempty"`
	// 收货状态
	TakeGoodsStatus string `json:"take_goods_status,omitempty" xml:"take_goods_status,omitempty"`
	// 理赔拒绝原因
	ClaimsResultDesc string `json:"claims_result_desc,omitempty" xml:"claims_result_desc,omitempty"`
	// 理赔币种
	ClaimCurrency string `json:"claim_currency,omitempty" xml:"claim_currency,omitempty"`
	// 赔付比例
	CompensationRatio string `json:"compensation_ratio,omitempty" xml:"compensation_ratio,omitempty"`
	// 包裹状态
	PackageStatus string `json:"package_status,omitempty" xml:"package_status,omitempty"`
	// 服务工单ID
	ServiceWorkOrderId int64 `json:"service_work_order_id,omitempty" xml:"service_work_order_id,omitempty"`
	// 主订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 理赔金额
	ClaimAmount int64 `json:"claim_amount,omitempty" xml:"claim_amount,omitempty"`
	// 预留扩展Map
	ExtensionMap *Extensionmap `json:"extension_map,omitempty" xml:"extension_map,omitempty"`
	// 理赔结果
	ClaimsResult bool `json:"claims_result,omitempty" xml:"claims_result,omitempty"`
}

var poolClaimsResultDto = sync.Pool{
	New: func() any {
		return new(ClaimsResultDto)
	},
}

// GetClaimsResultDto() 从对象池中获取ClaimsResultDto
func GetClaimsResultDto() *ClaimsResultDto {
	return poolClaimsResultDto.Get().(*ClaimsResultDto)
}

// ReleaseClaimsResultDto 释放ClaimsResultDto
func ReleaseClaimsResultDto(v *ClaimsResultDto) {
	v.ReportNo = ""
	v.TakeGoodsStatus = ""
	v.ClaimsResultDesc = ""
	v.ClaimCurrency = ""
	v.CompensationRatio = ""
	v.PackageStatus = ""
	v.ServiceWorkOrderId = 0
	v.OrderId = 0
	v.SubOrderId = 0
	v.ClaimAmount = 0
	v.ExtensionMap = nil
	v.ClaimsResult = false
	poolClaimsResultDto.Put(v)
}
