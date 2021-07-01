package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductReleaseAPIRequest
供应商铺货 API请求
tmall.supplychain.channel.product.release

供应商渠道铺货接口 */
type TmallSupplychainChannelProductReleaseAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// 渠道ID
	_channelCode int64
}

// New
