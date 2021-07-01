package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductIdEncryptAPIRequest
ICBU国际站商品加密接口 API请求
alibaba.icbu.product.id.encrypt

ICBU国际站，对混淆的产品ID加密。 */
type AlibabaIcbuProductIdEncryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 明文id
	_productId int64
}

// New
