package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyBrandSearchAPIRequest
星河-品牌搜索 API请求
alitrip.merchant.galaxy.brand.search

星河服务=获取雅高品牌信息 */
type AlitripMerchantGalaxyBrandSearchAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
}

// New
