package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainOrderSearchAPIRequest 获取火车票订单列表 API请求
// alitrip.btrip.train.order.search
//
// 第三方OA厂商获取自己的火车票数据
type AlitripBtripTrainOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripBtripTrainOrderSearchRequest 初始化AlitripBtripTrainOrderSearchAPIRequest对象
func NewAlitripBtripTrainOrderSearchRequest() *AlitripBtripTrainOrderSearchAPIRequest {
	return &AlitripBtripTrainOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripTrainOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.train.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripTrainOrderSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripBtripTrainOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripTrainOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
