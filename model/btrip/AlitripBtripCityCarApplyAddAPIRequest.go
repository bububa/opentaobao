package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCityCarApplyAddAPIRequest
三方市内用车申请单同步 API请求
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步 */
type AlitripBtripCityCarApplyAddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyAddRq
}

// NewAlitripBtripCityCarApplyAddRequest 初始化AlitripBtripCityCarApplyAddAPIRequest对象
func NewAlitripBtripCityCarApplyAddRequest() *AlitripBtripCityCarApplyAddAPIRequest {
	return &AlitripBtripCityCarApplyAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyAddAPIRequest) SetRq(_rq *CityCarApplyAddRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCityCarApplyAddAPIRequest) GetRq() *CityCarApplyAddRq {
	return r._rq
}
