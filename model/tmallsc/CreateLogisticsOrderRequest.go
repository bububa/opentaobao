package tmallsc

import (
	"sync"
)

// CreateLogisticsOrderRequest 结构体
type CreateLogisticsOrderRequest struct {
	// 外部单据id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// STAGE1:一阶段物流 STAGE2:二阶段物流
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 逗号分隔的工单id列表
	WorkcardIds string `json:"workcard_ids,omitempty" xml:"workcard_ids,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 快递公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 目的地四级地址编码
	ToAddressLocation string `json:"to_address_location,omitempty" xml:"to_address_location,omitempty"`
	// 目的地详细地址
	ToAddressDetail string `json:"to_address_detail,omitempty" xml:"to_address_detail,omitempty"`
	// 取件四级地址编码
	FromAddressLocation string `json:"from_address_location,omitempty" xml:"from_address_location,omitempty"`
	// 取件详细地址
	FromAddressDetail string `json:"from_address_detail,omitempty" xml:"from_address_detail,omitempty"`
}

var poolCreateLogisticsOrderRequest = sync.Pool{
	New: func() any {
		return new(CreateLogisticsOrderRequest)
	},
}

// GetCreateLogisticsOrderRequest() 从对象池中获取CreateLogisticsOrderRequest
func GetCreateLogisticsOrderRequest() *CreateLogisticsOrderRequest {
	return poolCreateLogisticsOrderRequest.Get().(*CreateLogisticsOrderRequest)
}

// ReleaseCreateLogisticsOrderRequest 释放CreateLogisticsOrderRequest
func ReleaseCreateLogisticsOrderRequest(v *CreateLogisticsOrderRequest) {
	v.OuterId = ""
	v.Type = ""
	v.WorkcardIds = ""
	v.MailNo = ""
	v.CompanyName = ""
	v.ToAddressLocation = ""
	v.ToAddressDetail = ""
	v.FromAddressLocation = ""
	v.FromAddressDetail = ""
	poolCreateLogisticsOrderRequest.Put(v)
}
