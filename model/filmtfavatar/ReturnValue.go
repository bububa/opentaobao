package filmtfavatar

import (
	"sync"
)

// ReturnValue 结构体
type ReturnValue struct {
	// 数据: 包含: 淘宝订单ID,系统商订单号,付款时间,核销时间,影院ID,影院名称,卖品名称,卖品内容,卖品来源,卖品结算单价,卖品数量,实际结算金额,卖品结算总价,影院卖品补贴,订单状态,退款状态,会员卡标识,备注,是否后结算,
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 数据数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolReturnValue = sync.Pool{
	New: func() any {
		return new(ReturnValue)
	},
}

// GetReturnValue() 从对象池中获取ReturnValue
func GetReturnValue() *ReturnValue {
	return poolReturnValue.Get().(*ReturnValue)
}

// ReleaseReturnValue 释放ReturnValue
func ReleaseReturnValue(v *ReturnValue) {
	v.Data = ""
	v.Count = 0
	poolReturnValue.Put(v)
}
