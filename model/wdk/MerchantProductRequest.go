package wdk

import (
	"sync"
)

// MerchantProductRequest 结构体
type MerchantProductRequest struct {
	// 商品条码
	Barcodes []string `json:"barcodes,omitempty" xml:"barcodes>string,omitempty"`
	// 类目属性信息
	CatProps []CatProps `json:"cat_props,omitempty" xml:"cat_props>cat_props,omitempty"`
	// 运输: 重量,单位: (g)
	TransportWeight string `json:"transport_weight,omitempty" xml:"transport_weight,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 税率
	TaxInvoice string `json:"tax_invoice,omitempty" xml:"tax_invoice,omitempty"`
	// 库存单位
	InventoryUnit string `json:"inventory_unit,omitempty" xml:"inventory_unit,omitempty"`
	// 价格,小数点2位
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 运输: 高度(mm)
	TransportHeight string `json:"transport_height,omitempty" xml:"transport_height,omitempty"`
	// 销售: 高度(mm)
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 销售: 长度(mm)
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 规格
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 销售: 重量(g)
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 运输: 宽度(mm)
	TransportWidth string `json:"transport_width,omitempty" xml:"transport_width,omitempty"`
	// 铺货渠道类目(key: 渠道编码, value: 渠道类目)
	MarketCategories string `json:"market_categories,omitempty" xml:"market_categories,omitempty"`
	// 运输: 长度(mm)
	TransportLength string `json:"transport_length,omitempty" xml:"transport_length,omitempty"`
	// 销售: 体积(cm^3)
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 销售: 宽度(mm)
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 均重
	AvgWeight string `json:"avg_weight,omitempty" xml:"avg_weight,omitempty"`
	// 优鲜类目ID
	YxCategoryId string `json:"yx_category_id,omitempty" xml:"yx_category_id,omitempty"`
	// 商家类目名路径
	RtCategoryNamePath string `json:"rt_category_name_path,omitempty" xml:"rt_category_name_path,omitempty"`
	// 运输: 体积,单位: (cm^3)
	TransportVolume string `json:"transport_volume,omitempty" xml:"transport_volume,omitempty"`
	// 优鲜类目名路径
	YxCategoryNamePath string `json:"yx_category_name_path,omitempty" xml:"yx_category_name_path,omitempty"`
	// 保质期(天)
	ExpirationDays int64 `json:"expiration_days,omitempty" xml:"expiration_days,omitempty"`
	// 商家类目ID
	RtCategoryId int64 `json:"rt_category_id,omitempty" xml:"rt_category_id,omitempty"`
	// 箱规数/包装数
	PackageNum int64 `json:"package_num,omitempty" xml:"package_num,omitempty"`
	// APP购买步长
	StepQuantity int64 `json:"step_quantity,omitempty" xml:"step_quantity,omitempty"`
	// 阿里标准叶子类目
	StandardCategoryId int64 `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty"`
	// 最小起订量
	PurchaseQuantity int64 `json:"purchase_quantity,omitempty" xml:"purchase_quantity,omitempty"`
	// RT商品ID
	RtItemCode int64 `json:"rt_item_code,omitempty" xml:"rt_item_code,omitempty"`
	// 效期方式(true:有效期,false:无效期)
	IsShelflife bool `json:"is_shelflife,omitempty" xml:"is_shelflife,omitempty"`
	// 是否称重
	IsWeight bool `json:"is_weight,omitempty" xml:"is_weight,omitempty"`
}

var poolMerchantProductRequest = sync.Pool{
	New: func() any {
		return new(MerchantProductRequest)
	},
}

// GetMerchantProductRequest() 从对象池中获取MerchantProductRequest
func GetMerchantProductRequest() *MerchantProductRequest {
	return poolMerchantProductRequest.Get().(*MerchantProductRequest)
}

// ReleaseMerchantProductRequest 释放MerchantProductRequest
func ReleaseMerchantProductRequest(v *MerchantProductRequest) {
	v.Barcodes = v.Barcodes[:0]
	v.CatProps = v.CatProps[:0]
	v.TransportWeight = ""
	v.Title = ""
	v.TaxInvoice = ""
	v.InventoryUnit = ""
	v.Price = ""
	v.TransportHeight = ""
	v.Height = ""
	v.Length = ""
	v.Specification = ""
	v.Weight = ""
	v.TransportWidth = ""
	v.MarketCategories = ""
	v.TransportLength = ""
	v.Volume = ""
	v.Width = ""
	v.AvgWeight = ""
	v.YxCategoryId = ""
	v.RtCategoryNamePath = ""
	v.TransportVolume = ""
	v.YxCategoryNamePath = ""
	v.ExpirationDays = 0
	v.RtCategoryId = 0
	v.PackageNum = 0
	v.StepQuantity = 0
	v.StandardCategoryId = 0
	v.PurchaseQuantity = 0
	v.RtItemCode = 0
	v.IsShelflife = false
	v.IsWeight = false
	poolMerchantProductRequest.Put(v)
}
