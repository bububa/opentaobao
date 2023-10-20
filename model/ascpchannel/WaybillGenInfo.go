package ascpchannel

import (
	"sync"
)

// WaybillGenInfo 结构体
type WaybillGenInfo struct {
	// 总体积
	TotalVolume string `json:"total_volume,omitempty" xml:"total_volume,omitempty"`
	// 外部单号
	RelatedOrderCode string `json:"related_order_code,omitempty" xml:"related_order_code,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 交易单号
	MainTradeNo string `json:"main_trade_no,omitempty" xml:"main_trade_no,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 拼单数量
	GatherNum int64 `json:"gather_num,omitempty" xml:"gather_num,omitempty"`
	// 是否拼单：0为否，1为是
	Gather int64 `json:"gather,omitempty" xml:"gather,omitempty"`
	// 总包件数
	TotalPackNum int64 `json:"total_pack_num,omitempty" xml:"total_pack_num,omitempty"`
	// 备注图片
	RemarkPicture *RemarkPicture `json:"remark_picture,omitempty" xml:"remark_picture,omitempty"`
	// 交易渠道：1为淘宝/天猫，2为其它，3为抖音，4为拼多多，5为京东，6为唯品会
	OrderChannel int64 `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
}

var poolWaybillGenInfo = sync.Pool{
	New: func() any {
		return new(WaybillGenInfo)
	},
}

// GetWaybillGenInfo() 从对象池中获取WaybillGenInfo
func GetWaybillGenInfo() *WaybillGenInfo {
	return poolWaybillGenInfo.Get().(*WaybillGenInfo)
}

// ReleaseWaybillGenInfo 释放WaybillGenInfo
func ReleaseWaybillGenInfo(v *WaybillGenInfo) {
	v.TotalVolume = ""
	v.RelatedOrderCode = ""
	v.Remark = ""
	v.MainTradeNo = ""
	v.SellerId = ""
	v.GatherNum = 0
	v.Gather = 0
	v.TotalPackNum = 0
	v.RemarkPicture = nil
	v.OrderChannel = 0
	poolWaybillGenInfo.Put(v)
}
