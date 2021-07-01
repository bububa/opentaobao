package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyShareGetAPIRequest
星河-获取小程序分享文案和图片 API请求
alitrip.merchant.galaxy.share.get

获取 雅高微信小程序分享素材文案和图片。 */
type AlitripMerchantGalaxyShareGetAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
}

// New
