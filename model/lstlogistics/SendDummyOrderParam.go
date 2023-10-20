package lstlogistics

import (
	"sync"
)

// SendDummyOrderParam 结构体
type SendDummyOrderParam struct {
	// 发货主订单列表
	MainOrderParamList []MainOrderParam `json:"main_order_param_list,omitempty" xml:"main_order_param_list>main_order_param,omitempty"`
	// 发货时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 备注
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

var poolSendDummyOrderParam = sync.Pool{
	New: func() any {
		return new(SendDummyOrderParam)
	},
}

// GetSendDummyOrderParam() 从对象池中获取SendDummyOrderParam
func GetSendDummyOrderParam() *SendDummyOrderParam {
	return poolSendDummyOrderParam.Get().(*SendDummyOrderParam)
}

// ReleaseSendDummyOrderParam 释放SendDummyOrderParam
func ReleaseSendDummyOrderParam(v *SendDummyOrderParam) {
	v.MainOrderParamList = v.MainOrderParamList[:0]
	v.SendTime = ""
	v.Remarks = ""
	poolSendDummyOrderParam.Put(v)
}
