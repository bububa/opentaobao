package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripTrainOrderSearchAPIRequest
获取火车票订单列表 API请求
alitrip.btrip.train.order.search

第三方OA厂商获取自己的火车票数据 */
type AlitripBtripTrainOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// New
