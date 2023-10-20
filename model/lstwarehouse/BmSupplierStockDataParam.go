package lstwarehouse

import (
	"sync"
)

// BmSupplierStockDataParam 结构体
type BmSupplierStockDataParam struct {
	// 供应商memberId
	SupplierMemberId string `json:"supplier_member_id,omitempty" xml:"supplier_member_id,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 数据日期，格式：yyyy-MM-dd
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolBmSupplierStockDataParam = sync.Pool{
	New: func() any {
		return new(BmSupplierStockDataParam)
	},
}

// GetBmSupplierStockDataParam() 从对象池中获取BmSupplierStockDataParam
func GetBmSupplierStockDataParam() *BmSupplierStockDataParam {
	return poolBmSupplierStockDataParam.Get().(*BmSupplierStockDataParam)
}

// ReleaseBmSupplierStockDataParam 释放BmSupplierStockDataParam
func ReleaseBmSupplierStockDataParam(v *BmSupplierStockDataParam) {
	v.SupplierMemberId = ""
	v.BrandName = ""
	v.StatDate = ""
	v.Page = 0
	v.Size = 0
	poolBmSupplierStockDataParam.Put(v)
}
