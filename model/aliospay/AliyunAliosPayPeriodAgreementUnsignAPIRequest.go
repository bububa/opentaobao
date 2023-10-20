package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayPeriodAgreementUnsignAPIRequest 周期扣款协议解约接口 API请求
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
type AliyunAliosPayPeriodAgreementUnsignAPIRequest struct {
	model.Params
	// 请求参数
	_periodAgreementUnsignRequest *PeriodAgreementUnsignRequest
}

// NewAliyunAliosPayPeriodAgreementUnsignRequest 初始化AliyunAliosPayPeriodAgreementUnsignAPIRequest对象
func NewAliyunAliosPayPeriodAgreementUnsignRequest() *AliyunAliosPayPeriodAgreementUnsignAPIRequest {
	return &AliyunAliosPayPeriodAgreementUnsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayPeriodAgreementUnsignAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.period.agreement.unsign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayPeriodAgreementUnsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayPeriodAgreementUnsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPeriodAgreementUnsignRequest is PeriodAgreementUnsignRequest Setter
// 请求参数
func (r *AliyunAliosPayPeriodAgreementUnsignAPIRequest) SetPeriodAgreementUnsignRequest(_periodAgreementUnsignRequest *PeriodAgreementUnsignRequest) error {
	r._periodAgreementUnsignRequest = _periodAgreementUnsignRequest
	r.Set("period_agreement_unsign_request", _periodAgreementUnsignRequest)
	return nil
}

// GetPeriodAgreementUnsignRequest PeriodAgreementUnsignRequest Getter
func (r AliyunAliosPayPeriodAgreementUnsignAPIRequest) GetPeriodAgreementUnsignRequest() *PeriodAgreementUnsignRequest {
	return r._periodAgreementUnsignRequest
}
