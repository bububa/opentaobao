package aliospay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayPeriodAgreementStatusGetAPIRequest 查询周期扣款签约状态 API请求
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
type AliyunAliosPayPeriodAgreementStatusGetAPIRequest struct {
	model.Params
	// 请求参数
	_getPeriodAgreementStatusRequest *GetPeriodAgreementStatusRequest
}

// NewAliyunAliosPayPeriodAgreementStatusGetRequest 初始化AliyunAliosPayPeriodAgreementStatusGetAPIRequest对象
func NewAliyunAliosPayPeriodAgreementStatusGetRequest() *AliyunAliosPayPeriodAgreementStatusGetAPIRequest {
	return &AliyunAliosPayPeriodAgreementStatusGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunAliosPayPeriodAgreementStatusGetAPIRequest) Reset() {
	r._getPeriodAgreementStatusRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayPeriodAgreementStatusGetAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.period.agreement.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayPeriodAgreementStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayPeriodAgreementStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetPeriodAgreementStatusRequest is GetPeriodAgreementStatusRequest Setter
// 请求参数
func (r *AliyunAliosPayPeriodAgreementStatusGetAPIRequest) SetGetPeriodAgreementStatusRequest(_getPeriodAgreementStatusRequest *GetPeriodAgreementStatusRequest) error {
	r._getPeriodAgreementStatusRequest = _getPeriodAgreementStatusRequest
	r.Set("get_period_agreement_status_request", _getPeriodAgreementStatusRequest)
	return nil
}

// GetGetPeriodAgreementStatusRequest GetPeriodAgreementStatusRequest Getter
func (r AliyunAliosPayPeriodAgreementStatusGetAPIRequest) GetGetPeriodAgreementStatusRequest() *GetPeriodAgreementStatusRequest {
	return r._getPeriodAgreementStatusRequest
}

var poolAliyunAliosPayPeriodAgreementStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunAliosPayPeriodAgreementStatusGetRequest()
	},
}

// GetAliyunAliosPayPeriodAgreementStatusGetRequest 从 sync.Pool 获取 AliyunAliosPayPeriodAgreementStatusGetAPIRequest
func GetAliyunAliosPayPeriodAgreementStatusGetAPIRequest() *AliyunAliosPayPeriodAgreementStatusGetAPIRequest {
	return poolAliyunAliosPayPeriodAgreementStatusGetAPIRequest.Get().(*AliyunAliosPayPeriodAgreementStatusGetAPIRequest)
}

// ReleaseAliyunAliosPayPeriodAgreementStatusGetAPIRequest 将 AliyunAliosPayPeriodAgreementStatusGetAPIRequest 放入 sync.Pool
func ReleaseAliyunAliosPayPeriodAgreementStatusGetAPIRequest(v *AliyunAliosPayPeriodAgreementStatusGetAPIRequest) {
	v.Reset()
	poolAliyunAliosPayPeriodAgreementStatusGetAPIRequest.Put(v)
}
