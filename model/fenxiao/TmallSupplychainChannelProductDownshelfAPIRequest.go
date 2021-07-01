package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductDownshelfAPIRequest
产品下架 API请求
tmall.supplychain.channel.product.downshelf

产品下架 */
type TmallSupplychainChannelProductDownshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// New
