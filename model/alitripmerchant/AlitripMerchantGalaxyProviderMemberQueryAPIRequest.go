package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyProviderMemberQueryAPIRequest
提供会员查询接口 API请求
alitrip.merchant.galaxy.provider.member.query

星河产品=提供会查询服务 */
type AlitripMerchantGalaxyProviderMemberQueryAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 查询参数
	_queryMemberParam *QueryMemberParam
}

// New
