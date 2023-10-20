package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbusinsurancerecommendAPIRequest 汽车票保险推荐接口 API请求
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
type AlitripbusinsurancerecommendAPIRequest struct {
	model.Params
	// 入参
	_insuranceRecommendRq *InsuranceRecommendRq
}

// NewAlitripbusinsurancerecommendRequest 初始化AlitripbusinsurancerecommendAPIRequest对象
func NewAlitripbusinsurancerecommendRequest() *AlitripbusinsurancerecommendAPIRequest {
	return &AlitripbusinsurancerecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbusinsurancerecommendAPIRequest) GetApiMethodName() string {
	return "alitrip.bus.insurance.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbusinsurancerecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbusinsurancerecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsuranceRecommendRq is InsuranceRecommendRq Setter
// 入参
func (r *AlitripbusinsurancerecommendAPIRequest) SetInsuranceRecommendRq(_insuranceRecommendRq *InsuranceRecommendRq) error {
	r._insuranceRecommendRq = _insuranceRecommendRq
	r.Set("insurance_recommend_rq", _insuranceRecommendRq)
	return nil
}

// GetInsuranceRecommendRq InsuranceRecommendRq Getter
func (r AlitripbusinsurancerecommendAPIRequest) GetInsuranceRecommendRq() *InsuranceRecommendRq {
	return r._insuranceRecommendRq
}
