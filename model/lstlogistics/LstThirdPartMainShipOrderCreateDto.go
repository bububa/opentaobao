package lstlogistics

import (
	"sync"
)

// LstThirdPartMainShipOrderCreateDto 结构体
type LstThirdPartMainShipOrderCreateDto struct {
	// 货品明细
	Details []LstThirdPartDetailShipOrderCreateDto `json:"details,omitempty" xml:"details>lst_third_part_detail_ship_order_create_dto,omitempty"`
	// 区
	ReceiverDistrict string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
	// 省
	ReceiverProvince string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 详细地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 店铺名称
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 订单创建时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 街道
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 邮编
	ReceiverZip string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
	// 外部订单ID，幂等key
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货人手机号
	ReceiverMobile int64 `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 订单实付金额，单位为分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}

var poolLstThirdPartMainShipOrderCreateDto = sync.Pool{
	New: func() any {
		return new(LstThirdPartMainShipOrderCreateDto)
	},
}

// GetLstThirdPartMainShipOrderCreateDto() 从对象池中获取LstThirdPartMainShipOrderCreateDto
func GetLstThirdPartMainShipOrderCreateDto() *LstThirdPartMainShipOrderCreateDto {
	return poolLstThirdPartMainShipOrderCreateDto.Get().(*LstThirdPartMainShipOrderCreateDto)
}

// ReleaseLstThirdPartMainShipOrderCreateDto 释放LstThirdPartMainShipOrderCreateDto
func ReleaseLstThirdPartMainShipOrderCreateDto(v *LstThirdPartMainShipOrderCreateDto) {
	v.Details = v.Details[:0]
	v.ReceiverDistrict = ""
	v.ReceiverProvince = ""
	v.BuyerMessage = ""
	v.ReceiverAddress = ""
	v.BuyerName = ""
	v.OrderCreateTime = ""
	v.ReceiverCity = ""
	v.ReceiverTown = ""
	v.ReceiverName = ""
	v.ReceiverZip = ""
	v.OutOrderId = ""
	v.ReceiverPhone = ""
	v.ReceiverMobile = 0
	v.PayFee = 0
	poolLstThirdPartMainShipOrderCreateDto.Put(v)
}
