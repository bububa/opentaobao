package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripFlightOrderSearchAPIRequest
获取机票订单列表 API请求
alitrip.btrip.flight.order.search

第三方OA厂商获取机票订单列表 */
type AlitripBtripFlightOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// New
