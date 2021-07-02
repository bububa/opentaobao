package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressOfferDsProductSimplequeryAPIResponse Dropshipper查询单个商品的简易信息 API返回值
// aliexpress.offer.ds.product.simplequery
//
// 提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
type AliexpressOfferDsProductSimplequeryAPIResponse struct {
	model.CommonResponse
	AliexpressOfferDsProductSimplequeryAPIResponseModel
}

// AliexpressOfferDsProductSimplequeryAPIResponseModel is Dropshipper查询单个商品的简易信息 成功返回结果
type AliexpressOfferDsProductSimplequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_offer_ds_product_simplequery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 网站折扣后价格
	ItemOfferSiteSalePrice string `json:"item_offer_site_sale_price,omitempty" xml:"item_offer_site_sale_price,omitempty"`
	// 库存
	TotalAvailableStock int64 `json:"total_available_stock,omitempty" xml:"total_available_stock,omitempty"`
	// 商品的SKU信息
	AeopAeProductSKUs []AeopAeProductSku `json:"aeop_ae_product_s_k_us,omitempty" xml:"aeop_ae_product_s_k_us>aeop_ae_product_sku,omitempty"`
	// 产品的状态
	ProductStatusType string `json:"product_status_type,omitempty" xml:"product_status_type,omitempty"`
}
