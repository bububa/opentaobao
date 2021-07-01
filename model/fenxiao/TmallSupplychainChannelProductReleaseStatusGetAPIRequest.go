package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductReleaseStatusGetAPIRequest
产品铺货状态查询 API请求
tmall.supplychain.channel.product.release.status.get

巴拿马战役渠道产品状态查询 */
type TmallSupplychainChannelProductReleaseStatusGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 渠道ID ( 台湾 : 111002 )
	_channelCode int64
}

// New
