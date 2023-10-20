package wdk

import (
	"sync"
)

// AlibabaWdkSkuScrollQueryModelList 结构体
type AlibabaWdkSkuScrollQueryModelList struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商品在机构内的生命周期，商品状态；A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡
	LifeStatus string `json:"life_status,omitempty" xml:"life_status,omitempty"`
	// 条码：多个条码用英文逗号分割
	Barcodes string `json:"barcodes,omitempty" xml:"barcodes,omitempty"`
	// 价格：单位元
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 会员价：单位元
	MemberPrice string `json:"member_price,omitempty" xml:"member_price,omitempty"`
	// 售卖单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 是否称重品：1称重品0非称重品
	WeightFlag string `json:"weight_flag,omitempty" xml:"weight_flag,omitempty"`
	// 商家类目编码
	MerchantCatCode string `json:"merchant_cat_code,omitempty" xml:"merchant_cat_code,omitempty"`
	// 子公司编码
	OrgNo string `json:"org_no,omitempty" xml:"org_no,omitempty"`
	// 门店编码
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 渠道店编码
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 渠道店类型：4淘宝
	ChannelCodes string `json:"channel_codes,omitempty" xml:"channel_codes,omitempty"`
	// 税收编码(查询返回使用)
	TaxClassNo string `json:"tax_class_no,omitempty" xml:"tax_class_no,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 默认供应商
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 销售规格描述
	SaleSpec string `json:"sale_spec,omitempty" xml:"sale_spec,omitempty"`
	// 后台平台类目编码
	BackCatCode string `json:"back_cat_code,omitempty" xml:"back_cat_code,omitempty"`
	// 进项税率
	InputTaxRate string `json:"input_tax_rate,omitempty" xml:"input_tax_rate,omitempty"`
	// 销项税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品牌编码
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 保质期天数，商品的保质期（单位：天）,0表示没有保质期
	ShelfLife string `json:"shelf_life,omitempty" xml:"shelf_life,omitempty"`
	// 业务类型 1：盒饭  2：生鲜
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 是否测试商品
	TestFlag int64 `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	// 是否服务商品
	ServiceFlag int64 `json:"service_flag,omitempty" xml:"service_flag,omitempty"`
	// 是否线上可售 是 0：否 1：是
	OnlineSaleFlag int64 `json:"online_sale_flag,omitempty" xml:"online_sale_flag,omitempty"`
	// 加工时间 单位：分钟
	ProcessingTime int64 `json:"processing_time,omitempty" xml:"processing_time,omitempty"`
}

var poolAlibabaWdkSkuScrollQueryModelList = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuScrollQueryModelList)
	},
}

// GetAlibabaWdkSkuScrollQueryModelList() 从对象池中获取AlibabaWdkSkuScrollQueryModelList
func GetAlibabaWdkSkuScrollQueryModelList() *AlibabaWdkSkuScrollQueryModelList {
	return poolAlibabaWdkSkuScrollQueryModelList.Get().(*AlibabaWdkSkuScrollQueryModelList)
}

// ReleaseAlibabaWdkSkuScrollQueryModelList 释放AlibabaWdkSkuScrollQueryModelList
func ReleaseAlibabaWdkSkuScrollQueryModelList(v *AlibabaWdkSkuScrollQueryModelList) {
	v.SkuCode = ""
	v.SkuName = ""
	v.LifeStatus = ""
	v.Barcodes = ""
	v.SalePrice = ""
	v.MemberPrice = ""
	v.SaleUnit = ""
	v.WeightFlag = ""
	v.MerchantCatCode = ""
	v.OrgNo = ""
	v.OuCode = ""
	v.ShopId = ""
	v.ChannelCodes = ""
	v.TaxClassNo = ""
	v.ModifiedTime = ""
	v.MerchantCode = ""
	v.SupplierNo = ""
	v.ShortTitle = ""
	v.SaleSpec = ""
	v.BackCatCode = ""
	v.InputTaxRate = ""
	v.TaxRate = ""
	v.BrandName = ""
	v.BrandCode = ""
	v.ShelfLife = ""
	v.BusinessType = 0
	v.TestFlag = 0
	v.ServiceFlag = 0
	v.OnlineSaleFlag = 0
	v.ProcessingTime = 0
	poolAlibabaWdkSkuScrollQueryModelList.Put(v)
}
