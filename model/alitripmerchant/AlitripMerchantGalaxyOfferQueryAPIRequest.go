package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOfferQueryAPIRequest
星河-offer查询 API请求
alitrip.merchant.galaxy.offer.query

根据offer的ID查询offer信息 */
type AlitripMerchantGalaxyOfferQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// offer活动ID
	_offerIds string
	// 渠道来源
	_offerChannel string
}

// New
