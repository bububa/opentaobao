package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripVehicleOrderSearchAPIRequest
用车订单查询接口 API请求
alitrip.btrip.vehicle.order.search

企业获取商旅用车订单数据 */
type AlitripBtripVehicleOrderSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// New
