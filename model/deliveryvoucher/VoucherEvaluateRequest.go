package deliveryvoucher

import (
	"sync"
)

// VoucherEvaluateRequest 结构体
type VoucherEvaluateRequest struct {
	// 券信息，券信息,最多20条券记录
	VoucherInfos []DeliveryVoucherInfoDto `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
	// 评价内容
	EvaluateContent string `json:"evaluate_content,omitempty" xml:"evaluate_content,omitempty"`
	// 可扩展字段
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 评价时间
	EvaluateTime string `json:"evaluate_time,omitempty" xml:"evaluate_time,omitempty"`
	// 操作时间
	OperateDate string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
	// 流水号:唯一，幂等性判断
	OpId string `json:"op_id,omitempty" xml:"op_id,omitempty"`
	// 第三方服务商标识,top appkey
	Provider string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 评价分数 1：失望；2：不满；3：一般；4：满意；5：惊喜
	EvaluateScore int64 `json:"evaluate_score,omitempty" xml:"evaluate_score,omitempty"`
	// 主交易订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolVoucherEvaluateRequest = sync.Pool{
	New: func() any {
		return new(VoucherEvaluateRequest)
	},
}

// GetVoucherEvaluateRequest() 从对象池中获取VoucherEvaluateRequest
func GetVoucherEvaluateRequest() *VoucherEvaluateRequest {
	return poolVoucherEvaluateRequest.Get().(*VoucherEvaluateRequest)
}

// ReleaseVoucherEvaluateRequest 释放VoucherEvaluateRequest
func ReleaseVoucherEvaluateRequest(v *VoucherEvaluateRequest) {
	v.VoucherInfos = v.VoucherInfos[:0]
	v.EvaluateContent = ""
	v.Extend = ""
	v.EvaluateTime = ""
	v.OperateDate = ""
	v.OpId = ""
	v.Provider = ""
	v.EvaluateScore = 0
	v.OrderId = 0
	poolVoucherEvaluateRequest.Put(v)
}
