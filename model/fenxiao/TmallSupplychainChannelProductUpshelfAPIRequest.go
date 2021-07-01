package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductUpshelfAPIRequest
产品上架 API请求
tmall.supplychain.channel.product.upshelf

上架渠道产品 */
type TmallSupplychainChannelProductUpshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// New
