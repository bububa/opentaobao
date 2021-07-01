package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressUsergrowthSearchItemsGetAPIRequest
第三方平台搜索AE商品 API请求
aliexpress.usergrowth.search.items.get

第三方平台的搜索服务   获取AE商品list */
type AliexpressUsergrowthSearchItemsGetAPIRequest struct {
	model.Params
	// user input keypods
	_keywords string
	// Third party tracking_id
	_trackingId string
	// currency code
	_currencyCode string
	// language
	_language string
	// page size
	_pageSize string
	// page index
	_pageIndex string
	// ship to country
	_countryCode string
}

// New
