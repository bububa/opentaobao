package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceProductSearchAPIRequest 搜索保险产品 API请求
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
type AlitripFlightInsuranceProductSearchAPIRequest struct {
	model.Params
	// 保险产品id
	_insurancePremiumId int64
}

// NewAlitripFlightInsuranceProductSearchRequest 初始化AlitripFlightInsuranceProductSearchAPIRequest对象
func NewAlitripFlightInsuranceProductSearchRequest() *AlitripFlightInsuranceProductSearchAPIRequest {
	return &AlitripFlightInsuranceProductSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightInsuranceProductSearchAPIRequest) Reset() {
	r._insurancePremiumId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightInsuranceProductSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.product.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightInsuranceProductSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightInsuranceProductSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsurancePremiumId is InsurancePremiumId Setter
// 保险产品id
func (r *AlitripFlightInsuranceProductSearchAPIRequest) SetInsurancePremiumId(_insurancePremiumId int64) error {
	r._insurancePremiumId = _insurancePremiumId
	r.Set("insurance_premium_id", _insurancePremiumId)
	return nil
}

// GetInsurancePremiumId InsurancePremiumId Getter
func (r AlitripFlightInsuranceProductSearchAPIRequest) GetInsurancePremiumId() int64 {
	return r._insurancePremiumId
}

var poolAlitripFlightInsuranceProductSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightInsuranceProductSearchRequest()
	},
}

// GetAlitripFlightInsuranceProductSearchRequest 从 sync.Pool 获取 AlitripFlightInsuranceProductSearchAPIRequest
func GetAlitripFlightInsuranceProductSearchAPIRequest() *AlitripFlightInsuranceProductSearchAPIRequest {
	return poolAlitripFlightInsuranceProductSearchAPIRequest.Get().(*AlitripFlightInsuranceProductSearchAPIRequest)
}

// ReleaseAlitripFlightInsuranceProductSearchAPIRequest 将 AlitripFlightInsuranceProductSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightInsuranceProductSearchAPIRequest(v *AlitripFlightInsuranceProductSearchAPIRequest) {
	v.Reset()
	poolAlitripFlightInsuranceProductSearchAPIRequest.Put(v)
}
