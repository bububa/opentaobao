package idle

import (
	"sync"
)

// TenderItemListVo 结构体
type TenderItemListVo struct {
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 场次（日期+编码）
	Schedule string `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// 等级
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 服务商仓库码
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 起拍价(分)
	StartPrice int64 `json:"start_price,omitempty" xml:"start_price,omitempty"`
	// 最终成交价(分)
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 订单信息
	OrderInfo *TenderOrderInfoVo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

var poolTenderItemListVo = sync.Pool{
	New: func() any {
		return new(TenderItemListVo)
	},
}

// GetTenderItemListVo() 从对象池中获取TenderItemListVo
func GetTenderItemListVo() *TenderItemListVo {
	return poolTenderItemListVo.Get().(*TenderItemListVo)
}

// ReleaseTenderItemListVo 释放TenderItemListVo
func ReleaseTenderItemListVo(v *TenderItemListVo) {
	v.ItemId = ""
	v.OrderId = ""
	v.ItemName = ""
	v.Status = ""
	v.Schedule = ""
	v.Degree = ""
	v.StartTime = ""
	v.EndTime = ""
	v.OutId = ""
	v.StartPrice = 0
	v.Price = 0
	v.OrderInfo = nil
	poolTenderItemListVo.Put(v)
}
