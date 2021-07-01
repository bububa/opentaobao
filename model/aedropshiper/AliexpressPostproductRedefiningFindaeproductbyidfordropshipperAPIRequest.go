package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest
Dropshipper查找商品信息接口 API请求
aliexpress.postproduct.redefining.findaeproductbyidfordropshipper

提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用 */
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest struct {
	model.Params
	// 商品ID
	_productId int64
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
}

// New
