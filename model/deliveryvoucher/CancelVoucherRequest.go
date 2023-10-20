package deliveryvoucher

import (
	"sync"
)

// CancelVoucherRequest 结构体
type CancelVoucherRequest struct {
	// 券信息,券信息,最多100条券记录
	VoucherInfos []DeliveryVoucherInfoDto `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
	// 操作时间
	OperateDate string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
	// 扩展参数
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 流水号:唯一，幂等性判断
	OpId string `json:"op_id,omitempty" xml:"op_id,omitempty"`
	// 第三方服务商标识,top appkey
	Provider string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 主订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolCancelVoucherRequest = sync.Pool{
	New: func() any {
		return new(CancelVoucherRequest)
	},
}

// GetCancelVoucherRequest() 从对象池中获取CancelVoucherRequest
func GetCancelVoucherRequest() *CancelVoucherRequest {
	return poolCancelVoucherRequest.Get().(*CancelVoucherRequest)
}

// ReleaseCancelVoucherRequest 释放CancelVoucherRequest
func ReleaseCancelVoucherRequest(v *CancelVoucherRequest) {
	v.VoucherInfos = v.VoucherInfos[:0]
	v.OperateDate = ""
	v.Extend = ""
	v.OpId = ""
	v.Provider = ""
	v.OrderId = 0
	poolCancelVoucherRequest.Put(v)
}
