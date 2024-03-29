package tmallnr

import (
	"sync"
)

// NrTimingOrderSoldQueryReqDto 结构体
type NrTimingOrderSoldQueryReqDto struct {
	// 要查询的订单创建结束时间，开始时间和结束时间之间最多相隔72小时
	EndCreated string `json:"end_created,omitempty" xml:"end_created,omitempty"`
	// 要查询的订单创建开始时间
	StartCreated string `json:"start_created,omitempty" xml:"start_created,omitempty"`
	// 业务标识，dss标识定时送业务；jsd表示极速达业务；qyls 区域零售业务标识
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 品牌商的sellerID
	BrandSellerId int64 `json:"brand_seller_id,omitempty" xml:"brand_seller_id,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小--当前限制了20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolNrTimingOrderSoldQueryReqDto = sync.Pool{
	New: func() any {
		return new(NrTimingOrderSoldQueryReqDto)
	},
}

// GetNrTimingOrderSoldQueryReqDto() 从对象池中获取NrTimingOrderSoldQueryReqDto
func GetNrTimingOrderSoldQueryReqDto() *NrTimingOrderSoldQueryReqDto {
	return poolNrTimingOrderSoldQueryReqDto.Get().(*NrTimingOrderSoldQueryReqDto)
}

// ReleaseNrTimingOrderSoldQueryReqDto 释放NrTimingOrderSoldQueryReqDto
func ReleaseNrTimingOrderSoldQueryReqDto(v *NrTimingOrderSoldQueryReqDto) {
	v.EndCreated = ""
	v.StartCreated = ""
	v.BizIdentity = ""
	v.BrandSellerId = 0
	v.PageNo = 0
	v.PageSize = 0
	poolNrTimingOrderSoldQueryReqDto.Put(v)
}
