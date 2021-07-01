package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductQuantityGetAPIRequest
渠道库存查询接口 API请求
tmall.supplychain.channel.product.quantity.get

渠道库存查询接口 */
type TmallSupplychainChannelProductQuantityGetAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// SKU ID
	_skuId int64
}

// New
