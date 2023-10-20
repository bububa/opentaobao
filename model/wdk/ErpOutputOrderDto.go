package wdk

import (
	"sync"
)

// ErpOutputOrderDto 结构体
type ErpOutputOrderDto struct {
	// 商品明细列表（子表）
	OutputItemInfos []OutputItemInfoDto `json:"output_item_infos,omitempty" xml:"output_item_infos>output_item_info_dto,omitempty"`
	// 出库时间，商家系统中记录的本批次商品的实际退货出库时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 单据号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 供应商code，可选值：当是退给供应商时，提供供应商编码；当是退给大仓时，提供大仓编码；当是调拨出库时，提供对方门店编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 店仓code，指的是出库对象，对应一个物理店或仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 单据类型，出库接口的单据类型包括退货单和调拨出单(1：退货单； 2：调拨出单)
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 单据子类型，出库单据类型为退货单时，需要进一步区分子类型为退供应商和退大仓（DC）(1：退供应商  2：退大仓 )
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}

var poolErpOutputOrderDto = sync.Pool{
	New: func() any {
		return new(ErpOutputOrderDto)
	},
}

// GetErpOutputOrderDto() 从对象池中获取ErpOutputOrderDto
func GetErpOutputOrderDto() *ErpOutputOrderDto {
	return poolErpOutputOrderDto.Get().(*ErpOutputOrderDto)
}

// ReleaseErpOutputOrderDto 释放ErpOutputOrderDto
func ReleaseErpOutputOrderDto(v *ErpOutputOrderDto) {
	v.OutputItemInfos = v.OutputItemInfos[:0]
	v.BizDate = ""
	v.BizOrderCode = ""
	v.Remark = ""
	v.SupplierCode = ""
	v.WarehouseCode = ""
	v.BizType = 0
	v.SubType = 0
	poolErpOutputOrderDto.Put(v)
}
