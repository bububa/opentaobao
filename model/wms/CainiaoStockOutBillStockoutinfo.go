package wms

import (
	"sync"
)

// CainiaoStockOutBillStockoutinfo 结构体
type CainiaoStockOutBillStockoutinfo struct {
	// 包裹信息列表，包含每个包裹使用的快递信息
	PackageInfoList []CainiaoStockOutBillPackageinfolist `json:"package_info_list,omitempty" xml:"package_info_list>cainiao_stock_out_bill_packageinfolist,omitempty"`
	// 订单商品列表
	OrderItemList []CainiaoStockOutBillOrderitemlist `json:"order_item_list,omitempty" xml:"order_item_list>cainiao_stock_out_bill_orderitemlist,omitempty"`
	// ERP订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓库订单编码，LBX号
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// 仓库订单完成时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 入库单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 单据类型 903 普通出库单 305 B2B出库单 901 退供出库单
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
}

var poolCainiaoStockOutBillStockoutinfo = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillStockoutinfo)
	},
}

// GetCainiaoStockOutBillStockoutinfo() 从对象池中获取CainiaoStockOutBillStockoutinfo
func GetCainiaoStockOutBillStockoutinfo() *CainiaoStockOutBillStockoutinfo {
	return poolCainiaoStockOutBillStockoutinfo.Get().(*CainiaoStockOutBillStockoutinfo)
}

// ReleaseCainiaoStockOutBillStockoutinfo 释放CainiaoStockOutBillStockoutinfo
func ReleaseCainiaoStockOutBillStockoutinfo(v *CainiaoStockOutBillStockoutinfo) {
	v.PackageInfoList = v.PackageInfoList[:0]
	v.OrderItemList = v.OrderItemList[:0]
	v.OrderCode = ""
	v.CnOrderCode = ""
	v.ConfirmTime = ""
	v.Status = ""
	v.OrderType = 0
	poolCainiaoStockOutBillStockoutinfo.Put(v)
}
