package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressUsergrowthRecommendItemsGetAPIRequest
第三方平台推荐商品 API请求
aliexpress.usergrowth.recommend.items.get

第三方平台的推荐AE商品   场景：skin 、底部推荐等 */
type AliexpressUsergrowthRecommendItemsGetAPIRequest struct {
	model.Params
	// third party trackingId
	_trackingId string
	// currency Code
	_currencyCode string
	// language
	_language string
	// user type
	_userTypeCode string
	// page index,start from 1
	_pageIndex string
	// page size
	_pageSize string
	// country code
	_countryCode string
}

// New
