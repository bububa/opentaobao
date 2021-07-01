package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyCityLikeAPIRequest
星河-酒店城市模糊查询 API请求
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息 */
type AlitripMerchantGalaxyCityLikeAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 城市模糊
	_cityName string
	// 0国内1国外
	_domestic int64
}

// New
