package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyWechatInfoAPIRequest
星河-获取微信用户的信息 API请求
alitrip.merchant.galaxy.wechat.info

获取微信用户的openId和unionId */
type AlitripMerchantGalaxyWechatInfoAPIRequest struct {
	model.Params
	// 租户的id
	_tenantKey string
	// 微信小程序获取的code
	_code string
}

// New
