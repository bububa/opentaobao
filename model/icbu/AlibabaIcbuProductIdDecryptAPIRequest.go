package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductIdDecryptAPIRequest
商品ID解密 API请求
alibaba.icbu.product.id.decrypt

对混淆的产品ID解密 */
type AlibabaIcbuProductIdDecryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 混淆后的商品ID
	_productId string
}

// New
