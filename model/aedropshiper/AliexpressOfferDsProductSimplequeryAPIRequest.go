package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressOfferDsProductSimplequeryAPIRequest
Dropshipper查询单个商品的简易信息 API请求
aliexpress.offer.ds.product.simplequery

提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用 */
type AliexpressOfferDsProductSimplequeryAPIRequest struct {
	model.Params
	// 商品ID
	_productId int64
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
}

// New
