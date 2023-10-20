package ascp

import (
	"sync"
)

// TmsOrderConfirmRequest 结构体
type TmsOrderConfirmRequest struct {
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 行业
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 服务商code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 当前物流节点
	LogisticStatus string `json:"logistic_status,omitempty" xml:"logistic_status,omitempty"`
	// 操作人
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 物流节点文本描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 物流节点发生变更时间的时间戳
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolTmsOrderConfirmRequest = sync.Pool{
	New: func() any {
		return new(TmsOrderConfirmRequest)
	},
}

// GetTmsOrderConfirmRequest() 从对象池中获取TmsOrderConfirmRequest
func GetTmsOrderConfirmRequest() *TmsOrderConfirmRequest {
	return poolTmsOrderConfirmRequest.Get().(*TmsOrderConfirmRequest)
}

// ReleaseTmsOrderConfirmRequest 释放TmsOrderConfirmRequest
func ReleaseTmsOrderConfirmRequest(v *TmsOrderConfirmRequest) {
	v.SellerId = ""
	v.ExpressCode = ""
	v.RequestId = ""
	v.Industry = ""
	v.CpCode = ""
	v.LogisticStatus = ""
	v.OperatorName = ""
	v.Description = ""
	v.Remark = ""
	v.Feature = ""
	v.OperateTime = 0
	v.RequestTime = 0
	poolTmsOrderConfirmRequest.Put(v)
}
