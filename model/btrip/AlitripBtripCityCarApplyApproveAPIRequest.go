package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCityCarApplyApproveAPIRequest
三方市内用车申请单审批 API请求
alitrip.btrip.city.car.apply.approve

三方市内用车申请单审批 */
type AlitripBtripCityCarApplyApproveAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyApproveRq
}

// NewAlitripBtripCityCarApplyApproveRequest 初始化AlitripBtripCityCarApplyApproveAPIRequest对象
func NewAlitripBtripCityCarApplyApproveRequest() *AlitripBtripCityCarApplyApproveAPIRequest {
	return &AlitripBtripCityCarApplyApproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyApproveAPIRequest) SetRq(_rq *CityCarApplyApproveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetRq() *CityCarApplyApproveRq {
	return r._rq
}
