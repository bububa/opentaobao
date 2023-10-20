package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest 报告解读订单-医生退款 API请求
// alibaba.alihealth.examination.report.diagnose.order.doctor.refund
//
// 报告解读订单医生退款
type AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest struct {
	model.Params
	// 退款入参
	_param *RefundForAikangDoctorRequest
}

// NewAlibabaalihealthexaminationreportdiagnoseorderdoctorrefundRequest 初始化AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseorderdoctorrefundRequest() *AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.doctor.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 退款入参
func (r *AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest) SetParam(_param *RefundForAikangDoctorRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest) GetParam() *RefundForAikangDoctorRequest {
	return r._param
}
