package wdk

import (
	"sync"
)

// ExpressPackage 结构体
type ExpressPackage struct {
	// 快递公司编码
	ExpressComCode string `json:"express_com_code,omitempty" xml:"express_com_code,omitempty"`
	// 快递公司名称
	ExpressComName string `json:"express_com_name,omitempty" xml:"express_com_name,omitempty"`
	// 运单号
	WayBillNo string `json:"way_bill_no,omitempty" xml:"way_bill_no,omitempty"`
	// 包裹ID
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 包裹中商品出库销售数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 包裹中商品出库库存数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// 包裹重量（单位g）
	PackageWeight string `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
}

var poolExpressPackage = sync.Pool{
	New: func() any {
		return new(ExpressPackage)
	},
}

// GetExpressPackage() 从对象池中获取ExpressPackage
func GetExpressPackage() *ExpressPackage {
	return poolExpressPackage.Get().(*ExpressPackage)
}

// ReleaseExpressPackage 释放ExpressPackage
func ReleaseExpressPackage(v *ExpressPackage) {
	v.ExpressComCode = ""
	v.ExpressComName = ""
	v.WayBillNo = ""
	v.PackageId = ""
	v.ActualSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.PackageWeight = ""
	poolExpressPackage.Put(v)
}
