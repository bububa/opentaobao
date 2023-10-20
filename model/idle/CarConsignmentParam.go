package idle

import (
	"sync"
)

// CarConsignmentParam 结构体
type CarConsignmentParam struct {
	// 订单id
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 主状态
	MainStatus string `json:"main_status,omitempty" xml:"main_status,omitempty"`
	// 子状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 服务商appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 额外属性信息
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 子状态
	SubStatus string `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
}

var poolCarConsignmentParam = sync.Pool{
	New: func() any {
		return new(CarConsignmentParam)
	},
}

// GetCarConsignmentParam() 从对象池中获取CarConsignmentParam
func GetCarConsignmentParam() *CarConsignmentParam {
	return poolCarConsignmentParam.Get().(*CarConsignmentParam)
}

// ReleaseCarConsignmentParam 释放CarConsignmentParam
func ReleaseCarConsignmentParam(v *CarConsignmentParam) {
	v.BizOrderId = ""
	v.MainStatus = ""
	v.OrderStatus = ""
	v.Appkey = ""
	v.Attribute = ""
	v.SubStatus = ""
	poolCarConsignmentParam.Put(v)
}
