package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriptrainordersearchAPIRequest 获取火车票订单列表 API请求
// alitrip.btrip.train.order.search
//
// 第三方OA厂商获取自己的火车票数据
type AlitripbtriptrainordersearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripbtriptrainordersearchRequest 初始化AlitripbtriptrainordersearchAPIRequest对象
func NewAlitripbtriptrainordersearchRequest() *AlitripbtriptrainordersearchAPIRequest {
	return &AlitripbtriptrainordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriptrainordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.train.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriptrainordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriptrainordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripbtriptrainordersearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtriptrainordersearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
