package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBusInsuranceRecommendAPIRequest 汽车票保险推荐接口 API请求
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
type AlitripBusInsuranceRecommendAPIRequest struct {
	model.Params
	// 入参
	_insuranceRecommendRq *InsuranceRecommendRq
}

// NewAlitripBusInsuranceRecommendRequest 初始化AlitripBusInsuranceRecommendAPIRequest对象
func NewAlitripBusInsuranceRecommendRequest() *AlitripBusInsuranceRecommendAPIRequest {
	return &AlitripBusInsuranceRecommendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBusInsuranceRecommendAPIRequest) Reset() {
	r._insuranceRecommendRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBusInsuranceRecommendAPIRequest) GetApiMethodName() string {
	return "alitrip.bus.insurance.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBusInsuranceRecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBusInsuranceRecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsuranceRecommendRq is InsuranceRecommendRq Setter
// 入参
func (r *AlitripBusInsuranceRecommendAPIRequest) SetInsuranceRecommendRq(_insuranceRecommendRq *InsuranceRecommendRq) error {
	r._insuranceRecommendRq = _insuranceRecommendRq
	r.Set("insurance_recommend_rq", _insuranceRecommendRq)
	return nil
}

// GetInsuranceRecommendRq InsuranceRecommendRq Getter
func (r AlitripBusInsuranceRecommendAPIRequest) GetInsuranceRecommendRq() *InsuranceRecommendRq {
	return r._insuranceRecommendRq
}

var poolAlitripBusInsuranceRecommendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBusInsuranceRecommendRequest()
	},
}

// GetAlitripBusInsuranceRecommendRequest 从 sync.Pool 获取 AlitripBusInsuranceRecommendAPIRequest
func GetAlitripBusInsuranceRecommendAPIRequest() *AlitripBusInsuranceRecommendAPIRequest {
	return poolAlitripBusInsuranceRecommendAPIRequest.Get().(*AlitripBusInsuranceRecommendAPIRequest)
}

// ReleaseAlitripBusInsuranceRecommendAPIRequest 将 AlitripBusInsuranceRecommendAPIRequest 放入 sync.Pool
func ReleaseAlitripBusInsuranceRecommendAPIRequest(v *AlitripBusInsuranceRecommendAPIRequest) {
	v.Reset()
	poolAlitripBusInsuranceRecommendAPIRequest.Put(v)
}
