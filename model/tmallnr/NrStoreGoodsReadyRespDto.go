package tmallnr

import (
	"sync"
)

// NrStoreGoodsReadyRespDto 结构体
type NrStoreGoodsReadyRespDto struct {
	// 取件码
	GotCode string `json:"got_code,omitempty" xml:"got_code,omitempty"`
	// 根据门店+sellerId按日期生成从1到N的数据
	ShortId string `json:"short_id,omitempty" xml:"short_id,omitempty"`
	// 核销码
	MaCode string `json:"ma_code,omitempty" xml:"ma_code,omitempty"`
	// 发货单号
	CompanyOrderNo string `json:"company_order_no,omitempty" xml:"company_order_no,omitempty"`
	// 发货公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 取消码
	CancelCode string `json:"cancel_code,omitempty" xml:"cancel_code,omitempty"`
	// 主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolNrStoreGoodsReadyRespDto = sync.Pool{
	New: func() any {
		return new(NrStoreGoodsReadyRespDto)
	},
}

// GetNrStoreGoodsReadyRespDto() 从对象池中获取NrStoreGoodsReadyRespDto
func GetNrStoreGoodsReadyRespDto() *NrStoreGoodsReadyRespDto {
	return poolNrStoreGoodsReadyRespDto.Get().(*NrStoreGoodsReadyRespDto)
}

// ReleaseNrStoreGoodsReadyRespDto 释放NrStoreGoodsReadyRespDto
func ReleaseNrStoreGoodsReadyRespDto(v *NrStoreGoodsReadyRespDto) {
	v.GotCode = ""
	v.ShortId = ""
	v.MaCode = ""
	v.CompanyOrderNo = ""
	v.CompanyName = ""
	v.CancelCode = ""
	v.MainOrderId = 0
	poolNrStoreGoodsReadyRespDto.Put(v)
}
