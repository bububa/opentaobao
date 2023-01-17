package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopHotelExceedapplyGetAPIRequest 商旅酒店第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.hotel.exceedapply.get
//
// 商旅酒店第三方超标审批单搜索接口
type AlitripBtripCorpopHotelExceedapplyGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopHotelExceedapplyGetRequest 初始化AlitripBtripCorpopHotelExceedapplyGetAPIRequest对象
func NewAlitripBtripCorpopHotelExceedapplyGetRequest() *AlitripBtripCorpopHotelExceedapplyGetAPIRequest {
	return &AlitripBtripCorpopHotelExceedapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.hotel.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopHotelExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
