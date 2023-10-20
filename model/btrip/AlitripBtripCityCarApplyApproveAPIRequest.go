package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCityCarApplyApproveAPIRequest 三方市内用车申请单审批 API请求
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
type AlitripBtripCityCarApplyApproveAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyApproveRq
}

// NewAlitripBtripCityCarApplyApproveRequest 初始化AlitripBtripCityCarApplyApproveAPIRequest对象
func NewAlitripBtripCityCarApplyApproveRequest() *AlitripBtripCityCarApplyApproveAPIRequest {
	return &AlitripBtripCityCarApplyApproveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCityCarApplyApproveAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyApproveAPIRequest) SetRq(_rq *CityCarApplyApproveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCityCarApplyApproveAPIRequest) GetRq() *CityCarApplyApproveRq {
	return r._rq
}

var poolAlitripBtripCityCarApplyApproveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCityCarApplyApproveRequest()
	},
}

// GetAlitripBtripCityCarApplyApproveRequest 从 sync.Pool 获取 AlitripBtripCityCarApplyApproveAPIRequest
func GetAlitripBtripCityCarApplyApproveAPIRequest() *AlitripBtripCityCarApplyApproveAPIRequest {
	return poolAlitripBtripCityCarApplyApproveAPIRequest.Get().(*AlitripBtripCityCarApplyApproveAPIRequest)
}

// ReleaseAlitripBtripCityCarApplyApproveAPIRequest 将 AlitripBtripCityCarApplyApproveAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCityCarApplyApproveAPIRequest(v *AlitripBtripCityCarApplyApproveAPIRequest) {
	v.Reset()
	poolAlitripBtripCityCarApplyApproveAPIRequest.Put(v)
}
