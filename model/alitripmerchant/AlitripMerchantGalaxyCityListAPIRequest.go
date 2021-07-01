package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyCityListAPIRequest
星河-酒店城市列表展示 API请求
alitrip.merchant.galaxy.city.list

雅高酒店城市列表展示，并且首字母列出酒店城市 */
type AlitripMerchantGalaxyCityListAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 0国内 1国外
	_domestic int64
}

// New
