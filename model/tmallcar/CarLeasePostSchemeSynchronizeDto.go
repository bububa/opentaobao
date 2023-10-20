package tmallcar

import (
	"sync"
)

// CarLeasePostSchemeSynchronizeDto 结构体
type CarLeasePostSchemeSynchronizeDto struct {
	// 租后方案
	SchemeAfterLeaseList []CarLeasePostSchemeDto `json:"scheme_after_lease_list,omitempty" xml:"scheme_after_lease_list>car_lease_post_scheme_dto,omitempty"`
	// 合同到期时间
	ContractEndDate string `json:"contract_end_date,omitempty" xml:"contract_end_date,omitempty"`
	// 不能使用原因code
	ReasonCode string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	// 不能使用原因描述
	ReasonDesc string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
	// 0:不能使用,1:可以使用
	CanSelect int64 `json:"can_select,omitempty" xml:"can_select,omitempty"`
	// 最后一期是否已经支付,0未支付,1已支付
	LastPayStatus int64 `json:"last_pay_status,omitempty" xml:"last_pay_status,omitempty"`
	// 天猫开新车订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolCarLeasePostSchemeSynchronizeDto = sync.Pool{
	New: func() any {
		return new(CarLeasePostSchemeSynchronizeDto)
	},
}

// GetCarLeasePostSchemeSynchronizeDto() 从对象池中获取CarLeasePostSchemeSynchronizeDto
func GetCarLeasePostSchemeSynchronizeDto() *CarLeasePostSchemeSynchronizeDto {
	return poolCarLeasePostSchemeSynchronizeDto.Get().(*CarLeasePostSchemeSynchronizeDto)
}

// ReleaseCarLeasePostSchemeSynchronizeDto 释放CarLeasePostSchemeSynchronizeDto
func ReleaseCarLeasePostSchemeSynchronizeDto(v *CarLeasePostSchemeSynchronizeDto) {
	v.SchemeAfterLeaseList = v.SchemeAfterLeaseList[:0]
	v.ContractEndDate = ""
	v.ReasonCode = ""
	v.ReasonDesc = ""
	v.CanSelect = 0
	v.LastPayStatus = 0
	v.OrderId = 0
	poolCarLeasePostSchemeSynchronizeDto.Put(v)
}
