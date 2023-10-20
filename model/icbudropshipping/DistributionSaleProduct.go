package icbudropshipping

import (
	"sync"
)

// DistributionSaleProduct 结构体
type DistributionSaleProduct struct {
	// product keywords
	Keywords []string `json:"keywords,omitempty" xml:"keywords>string,omitempty"`
	// Ladder delivery List
	LadderPeriodList []LadderPeriod `json:"ladder_period_list,omitempty" xml:"ladder_period_list>ladder_period,omitempty"`
	// product sku list
	ProductSkuList []ProductSku `json:"product_sku_list,omitempty" xml:"product_sku_list>product_sku,omitempty"`
	// product image list
	ImageUrlList []string `json:"image_url_list,omitempty" xml:"image_url_list>string,omitempty"`
	// product description html，&lt;br /&gt; It is the transferred string, the applicable party needs to reverse the string
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// product detail Url
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// main image url
	MainImageUrl string `json:"main_image_url,omitempty" xml:"main_image_url,omitempty"`
	// product name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// price Range
	PriceRange string `json:"price_range,omitempty" xml:"price_range,omitempty"`
	// Store Id
	ECompanyId string `json:"e_company_id,omitempty" xml:"e_company_id,omitempty"`
	// the top category name
	TopCategoryName string `json:"top_category_name,omitempty" xml:"top_category_name,omitempty"`
	// the leaf category name
	LeafCategoryName string `json:"leaf_category_name,omitempty" xml:"leaf_category_name,omitempty"`
	// min order quantity and Price
	MoqAndPrice *MoqAndPrice `json:"moq_and_price,omitempty" xml:"moq_and_price,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// Determine whether this product can be ordered
	IsCanPlaceOrder bool `json:"is_can_place_order,omitempty" xml:"is_can_place_order,omitempty"`
}

var poolDistributionSaleProduct = sync.Pool{
	New: func() any {
		return new(DistributionSaleProduct)
	},
}

// GetDistributionSaleProduct() 从对象池中获取DistributionSaleProduct
func GetDistributionSaleProduct() *DistributionSaleProduct {
	return poolDistributionSaleProduct.Get().(*DistributionSaleProduct)
}

// ReleaseDistributionSaleProduct 释放DistributionSaleProduct
func ReleaseDistributionSaleProduct(v *DistributionSaleProduct) {
	v.Keywords = v.Keywords[:0]
	v.LadderPeriodList = v.LadderPeriodList[:0]
	v.ProductSkuList = v.ProductSkuList[:0]
	v.ImageUrlList = v.ImageUrlList[:0]
	v.Description = ""
	v.DetailUrl = ""
	v.MainImageUrl = ""
	v.Name = ""
	v.PriceRange = ""
	v.ECompanyId = ""
	v.TopCategoryName = ""
	v.LeafCategoryName = ""
	v.MoqAndPrice = nil
	v.ProductId = 0
	v.IsCanPlaceOrder = false
	poolDistributionSaleProduct.Put(v)
}
