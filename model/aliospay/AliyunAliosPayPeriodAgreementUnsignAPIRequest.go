package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayperiodagreementunsignAPIRequest 周期扣款协议解约接口 API请求
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
type AliyunaliospayperiodagreementunsignAPIRequest struct {
	model.Params
	// 请求参数
	_periodAgreementUnsignRequest *PeriodAgreementUnsignRequest
}

// NewAliyunaliospayperiodagreementunsignRequest 初始化AliyunaliospayperiodagreementunsignAPIRequest对象
func NewAliyunaliospayperiodagreementunsignRequest() *AliyunaliospayperiodagreementunsignAPIRequest {
	return &AliyunaliospayperiodagreementunsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunaliospayperiodagreementunsignAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.period.agreement.unsign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunaliospayperiodagreementunsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunaliospayperiodagreementunsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPeriodAgreementUnsignRequest is PeriodAgreementUnsignRequest Setter
// 请求参数
func (r *AliyunaliospayperiodagreementunsignAPIRequest) SetPeriodAgreementUnsignRequest(_periodAgreementUnsignRequest *PeriodAgreementUnsignRequest) error {
	r._periodAgreementUnsignRequest = _periodAgreementUnsignRequest
	r.Set("period_agreement_unsign_request", _periodAgreementUnsignRequest)
	return nil
}

// GetPeriodAgreementUnsignRequest PeriodAgreementUnsignRequest Getter
func (r AliyunaliospayperiodagreementunsignAPIRequest) GetPeriodAgreementUnsignRequest() *PeriodAgreementUnsignRequest {
	return r._periodAgreementUnsignRequest
}
