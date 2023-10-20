package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceproductsearchAPIRequest 搜索保险产品 API请求
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
type AlitripflightinsuranceproductsearchAPIRequest struct {
	model.Params
	// 保险产品id
	_insurancePremiumId int64
}

// NewAlitripflightinsuranceproductsearchRequest 初始化AlitripflightinsuranceproductsearchAPIRequest对象
func NewAlitripflightinsuranceproductsearchRequest() *AlitripflightinsuranceproductsearchAPIRequest {
	return &AlitripflightinsuranceproductsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightinsuranceproductsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.product.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightinsuranceproductsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightinsuranceproductsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsurancePremiumId is InsurancePremiumId Setter
// 保险产品id
func (r *AlitripflightinsuranceproductsearchAPIRequest) SetInsurancePremiumId(_insurancePremiumId int64) error {
	r._insurancePremiumId = _insurancePremiumId
	r.Set("insurance_premium_id", _insurancePremiumId)
	return nil
}

// GetInsurancePremiumId InsurancePremiumId Getter
func (r AlitripflightinsuranceproductsearchAPIRequest) GetInsurancePremiumId() int64 {
	return r._insurancePremiumId
}
