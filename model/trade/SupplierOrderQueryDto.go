package trade

import (
	"sync"
)

// SupplierOrderQueryDto 结构体
type SupplierOrderQueryDto struct {
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 查询交易创建时间在2020-06-03的订单，只接受yyyy-MM-dd格式的字符串
	TradeCreateDate string `json:"trade_create_date,omitempty" xml:"trade_create_date,omitempty"`
	// 供货商身份标识，比如大润发就传DRF
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// 分页游标
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页参数，页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolSupplierOrderQueryDto = sync.Pool{
	New: func() any {
		return new(SupplierOrderQueryDto)
	},
}

// GetSupplierOrderQueryDto() 从对象池中获取SupplierOrderQueryDto
func GetSupplierOrderQueryDto() *SupplierOrderQueryDto {
	return poolSupplierOrderQueryDto.Get().(*SupplierOrderQueryDto)
}

// ReleaseSupplierOrderQueryDto 释放SupplierOrderQueryDto
func ReleaseSupplierOrderQueryDto(v *SupplierOrderQueryDto) {
	v.OuterStoreId = ""
	v.TradeCreateDate = ""
	v.Supplier = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolSupplierOrderQueryDto.Put(v)
}
