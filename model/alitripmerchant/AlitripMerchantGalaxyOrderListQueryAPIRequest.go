package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderListQueryAPIRequest
星河-订单列表查询 API请求
alitrip.merchant.galaxy.order.list.query

为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期 */
type AlitripMerchantGalaxyOrderListQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户token
	_token string
	// 订单状态
	_orderStatus string
	// 入住时间
	_startTime string
	// 入住时间
	_endTime string
	// 页数
	_page int64
	// 每页行数
	_row int64
}

// New
