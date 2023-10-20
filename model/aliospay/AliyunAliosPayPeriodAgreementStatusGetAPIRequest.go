package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayperiodagreementstatusgetAPIRequest 查询周期扣款签约状态 API请求
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
type AliyunaliospayperiodagreementstatusgetAPIRequest struct {
	model.Params
	// 请求参数
	_getPeriodAgreementStatusRequest *GetPeriodAgreementStatusRequest
}

// NewAliyunaliospayperiodagreementstatusgetRequest 初始化AliyunaliospayperiodagreementstatusgetAPIRequest对象
func NewAliyunaliospayperiodagreementstatusgetRequest() *AliyunaliospayperiodagreementstatusgetAPIRequest {
	return &AliyunaliospayperiodagreementstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunaliospayperiodagreementstatusgetAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.period.agreement.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunaliospayperiodagreementstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunaliospayperiodagreementstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetPeriodAgreementStatusRequest is GetPeriodAgreementStatusRequest Setter
// 请求参数
func (r *AliyunaliospayperiodagreementstatusgetAPIRequest) SetGetPeriodAgreementStatusRequest(_getPeriodAgreementStatusRequest *GetPeriodAgreementStatusRequest) error {
	r._getPeriodAgreementStatusRequest = _getPeriodAgreementStatusRequest
	r.Set("get_period_agreement_status_request", _getPeriodAgreementStatusRequest)
	return nil
}

// GetGetPeriodAgreementStatusRequest GetPeriodAgreementStatusRequest Getter
func (r AliyunaliospayperiodagreementstatusgetAPIRequest) GetGetPeriodAgreementStatusRequest() *GetPeriodAgreementStatusRequest {
	return r._getPeriodAgreementStatusRequest
}
