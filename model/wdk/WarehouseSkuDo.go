package wdk

import (
	"sync"
)

// WarehouseSkuDo 结构体
type WarehouseSkuDo struct {
	// 商品条码
	Barcodes []WarehouseSkuBarcodeDo `json:"barcodes,omitempty" xml:"barcodes>warehouse_sku_barcode_do,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 含量
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 配货规格
	DeliverySpec string `json:"delivery_spec,omitempty" xml:"delivery_spec,omitempty"`
	// 配货单位
	DeliveryUnit string `json:"delivery_unit,omitempty" xml:"delivery_unit,omitempty"`
	// 配送方式，1-统配、2-直配、3-越库
	DeliveryWay string `json:"delivery_way,omitempty" xml:"delivery_way,omitempty"`
	// 新建时间
	GmtCreateTime string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
	// 进项税率
	InputTaxRate string `json:"input_tax_rate,omitempty" xml:"input_tax_rate,omitempty"`
	// 库存单位
	InventoryUnit string `json:"inventory_unit,omitempty" xml:"inventory_unit,omitempty"`
	// 商品状态，A-正常、T-暂时停采、C-淘汰出清、R-清退、L-季节性商品休眠、D-删除封挡、E-停售(紧急下架)、U-未启用（只是建档，还未进货）
	LifeStatus string `json:"life_status,omitempty" xml:"life_status,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 超收比例
	OverloadRate string `json:"overload_rate,omitempty" xml:"overload_rate,omitempty"`
	// 厂商名称
	ProducerName string `json:"producer_name,omitempty" xml:"producer_name,omitempty"`
	// 产地，多个产地使用逗号分割
	ProducerPlace string `json:"producer_place,omitempty" xml:"producer_place,omitempty"`
	// 采购规格
	PurchaseSpec string `json:"purchase_spec,omitempty" xml:"purchase_spec,omitempty"`
	// 采购单位
	PurchaseUnit string `json:"purchase_unit,omitempty" xml:"purchase_unit,omitempty"`
	// 简称
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 存储方式，241-常温、242-冷藏、243-冷冻、635-热链、636-室温、637-鲜活
	Storage string `json:"storage,omitempty" xml:"storage,omitempty"`
	// 当前供应商编码
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 进项税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 禁收时限
	ForbidReceiveDays int64 `json:"forbid_receive_days,omitempty" xml:"forbid_receive_days,omitempty"`
	// 禁售时限
	ForbidSalesDays int64 `json:"forbid_sales_days,omitempty" xml:"forbid_sales_days,omitempty"`
	// 商品类目
	MerchantCatId int64 `json:"merchant_cat_id,omitempty" xml:"merchant_cat_id,omitempty"`
	// 保质期天数
	Period int64 `json:"period,omitempty" xml:"period,omitempty"`
	// 保质期预警天数
	WarnDays int64 `json:"warn_days,omitempty" xml:"warn_days,omitempty"`
	// 是否进口商品
	ImportFlag bool `json:"import_flag,omitempty" xml:"import_flag,omitempty"`
	// 是否称重商品
	WeightFlag bool `json:"weight_flag,omitempty" xml:"weight_flag,omitempty"`
}

var poolWarehouseSkuDo = sync.Pool{
	New: func() any {
		return new(WarehouseSkuDo)
	},
}

// GetWarehouseSkuDo() 从对象池中获取WarehouseSkuDo
func GetWarehouseSkuDo() *WarehouseSkuDo {
	return poolWarehouseSkuDo.Get().(*WarehouseSkuDo)
}

// ReleaseWarehouseSkuDo 释放WarehouseSkuDo
func ReleaseWarehouseSkuDo(v *WarehouseSkuDo) {
	v.Barcodes = v.Barcodes[:0]
	v.BrandName = ""
	v.Content = ""
	v.DeliverySpec = ""
	v.DeliveryUnit = ""
	v.DeliveryWay = ""
	v.GmtCreateTime = ""
	v.InputTaxRate = ""
	v.InventoryUnit = ""
	v.LifeStatus = ""
	v.MerchantCode = ""
	v.OverloadRate = ""
	v.ProducerName = ""
	v.ProducerPlace = ""
	v.PurchaseSpec = ""
	v.PurchaseUnit = ""
	v.ShortTitle = ""
	v.SkuCode = ""
	v.SkuName = ""
	v.Storage = ""
	v.SupplierNo = ""
	v.TaxRate = ""
	v.WarehouseCode = ""
	v.ForbidReceiveDays = 0
	v.ForbidSalesDays = 0
	v.MerchantCatId = 0
	v.Period = 0
	v.WarnDays = 0
	v.ImportFlag = false
	v.WeightFlag = false
	poolWarehouseSkuDo.Put(v)
}
