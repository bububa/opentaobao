package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphotelordersearchAPIRequest 搜索酒店订单列表 API请求
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
type AlitripbtriphotelordersearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripbtriphotelordersearchRequest 初始化AlitripbtriphotelordersearchAPIRequest对象
func NewAlitripbtriphotelordersearchRequest() *AlitripbtriphotelordersearchAPIRequest {
	return &AlitripbtriphotelordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphotelordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphotelordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphotelordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripbtriphotelordersearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtriphotelordersearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
