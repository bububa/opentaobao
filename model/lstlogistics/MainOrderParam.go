package lstlogistics

import (
	"sync"
)

// MainOrderParam 结构体
type MainOrderParam struct {
	// 发货子订单列表
	SubOrderParamList []SubOrderParam `json:"sub_order_param_list,omitempty" xml:"sub_order_param_list>sub_order_param,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolMainOrderParam = sync.Pool{
	New: func() any {
		return new(MainOrderParam)
	},
}

// GetMainOrderParam() 从对象池中获取MainOrderParam
func GetMainOrderParam() *MainOrderParam {
	return poolMainOrderParam.Get().(*MainOrderParam)
}

// ReleaseMainOrderParam 释放MainOrderParam
func ReleaseMainOrderParam(v *MainOrderParam) {
	v.SubOrderParamList = v.SubOrderParamList[:0]
	v.MainOrderId = 0
	poolMainOrderParam.Put(v)
}
