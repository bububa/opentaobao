package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayPeriodAgreementPayAPIRequest 周期扣款支付接口 API请求
// aliyun.alios.pay.period.agreement.pay
//
// 周期扣款支付接口，商户服务端通过该接口完成后续期数的支付
type AliyunAliosPayPeriodAgreementPayAPIRequest struct {
	model.Params
	// 请求参数
	_periodAgreementPayRequest *PeriodAgreementPayRequest
}

// NewAliyunAliosPayPeriodAgreementPayRequest 初始化AliyunAliosPayPeriodAgreementPayAPIRequest对象
func NewAliyunAliosPayPeriodAgreementPayRequest() *AliyunAliosPayPeriodAgreementPayAPIRequest {
	return &AliyunAliosPayPeriodAgreementPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayPeriodAgreementPayAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.period.agreement.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayPeriodAgreementPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayPeriodAgreementPayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPeriodAgreementPayRequest is PeriodAgreementPayRequest Setter
// 请求参数
func (r *AliyunAliosPayPeriodAgreementPayAPIRequest) SetPeriodAgreementPayRequest(_periodAgreementPayRequest *PeriodAgreementPayRequest) error {
	r._periodAgreementPayRequest = _periodAgreementPayRequest
	r.Set("period_agreement_pay_request", _periodAgreementPayRequest)
	return nil
}

// GetPeriodAgreementPayRequest PeriodAgreementPayRequest Getter
func (r AliyunAliosPayPeriodAgreementPayAPIRequest) GetPeriodAgreementPayRequest() *PeriodAgreementPayRequest {
	return r._periodAgreementPayRequest
}
