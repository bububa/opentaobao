package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelOrderSearchAPIRequest
搜索酒店订单列表 API请求
alitrip.btrip.hotel.order.search

企业获取商旅酒店订单数据 */
type AlitripBtripHotelOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// New
