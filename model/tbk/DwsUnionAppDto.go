package tbk

import (
	"sync"
)

// DwsUnionAppDto 结构体
type DwsUnionAppDto struct {
	// 付款预估收入。指买家付款金额为基数，预估您可能获得的收入。因买家退款等原因，可能与结算预估收入不一致
	PrePubShareFee string `json:"pre_pub_share_fee,omitempty" xml:"pre_pub_share_fee,omitempty"`
	// 支付金额
	AlipayAmt string `json:"alipay_amt,omitempty" xml:"alipay_amt,omitempty"`
	// 人群内成交人数
	AlipayUserNumInUcrowd int64 `json:"alipay_user_num_in_ucrowd,omitempty" xml:"alipay_user_num_in_ucrowd,omitempty"`
	// 支付人数
	AlipayUserNum int64 `json:"alipay_user_num,omitempty" xml:"alipay_user_num,omitempty"`
	// 支付订单数
	AlipayNum int64 `json:"alipay_num,omitempty" xml:"alipay_num,omitempty"`
	// 点击人数
	ClickUvUnion int64 `json:"click_uv_union,omitempty" xml:"click_uv_union,omitempty"`
	// 点击数
	ClickPvUnion int64 `json:"click_pv_union,omitempty" xml:"click_pv_union,omitempty"`
}

var poolDwsUnionAppDto = sync.Pool{
	New: func() any {
		return new(DwsUnionAppDto)
	},
}

// GetDwsUnionAppDto() 从对象池中获取DwsUnionAppDto
func GetDwsUnionAppDto() *DwsUnionAppDto {
	return poolDwsUnionAppDto.Get().(*DwsUnionAppDto)
}

// ReleaseDwsUnionAppDto 释放DwsUnionAppDto
func ReleaseDwsUnionAppDto(v *DwsUnionAppDto) {
	v.PrePubShareFee = ""
	v.AlipayAmt = ""
	v.AlipayUserNumInUcrowd = 0
	v.AlipayUserNum = 0
	v.AlipayNum = 0
	v.ClickUvUnion = 0
	v.ClickPvUnion = 0
	poolDwsUnionAppDto.Put(v)
}
