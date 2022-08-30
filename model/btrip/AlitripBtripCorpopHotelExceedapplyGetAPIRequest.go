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
	_rq *OpenIsvSearchRQ
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
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopHotelExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRQ) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopHotelExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRQ {
	return r._rq
}
