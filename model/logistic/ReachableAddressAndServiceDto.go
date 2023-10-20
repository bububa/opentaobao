package logistic

import (
	"sync"
)

// ReachableAddressAndServiceDto 结构体
type ReachableAddressAndServiceDto struct {
	// 服务列表,每一项必须为json的string格式,快运必填,快递为空则默认为&#39;标准快递&#39;
	ServiceCodeList []string `json:"service_code_list,omitempty" xml:"service_code_list>string,omitempty"`
	// 每条收发地址的key,用户自定义,每次请求多个地址不能重复
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 淘宝开放地址ID
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// C2M&amp;1688开放地址ID
	Caid string `json:"caid,omitempty" xml:"caid,omitempty"`
	// 收货地址
	ReceiveAddress *ReceiveAddress `json:"receive_address,omitempty" xml:"receive_address,omitempty"`
	// 发货地址
	SendAddress *AddressDto `json:"send_address,omitempty" xml:"send_address,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolReachableAddressAndServiceDto = sync.Pool{
	New: func() any {
		return new(ReachableAddressAndServiceDto)
	},
}

// GetReachableAddressAndServiceDto() 从对象池中获取ReachableAddressAndServiceDto
func GetReachableAddressAndServiceDto() *ReachableAddressAndServiceDto {
	return poolReachableAddressAndServiceDto.Get().(*ReachableAddressAndServiceDto)
}

// ReleaseReachableAddressAndServiceDto 释放ReachableAddressAndServiceDto
func ReleaseReachableAddressAndServiceDto(v *ReachableAddressAndServiceDto) {
	v.ServiceCodeList = v.ServiceCodeList[:0]
	v.ObjectId = ""
	v.Oaid = ""
	v.Caid = ""
	v.ReceiveAddress = nil
	v.SendAddress = nil
	v.OrderId = 0
	poolReachableAddressAndServiceDto.Put(v)
}
