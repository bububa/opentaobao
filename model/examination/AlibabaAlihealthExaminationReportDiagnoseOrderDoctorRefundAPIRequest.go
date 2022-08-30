package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest 报告解读订单-医生退款 API请求
// alibaba.alihealth.examination.report.diagnose.order.doctor.refund
//
// 报告解读订单医生退款
type AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest struct {
	model.Params
	// 退款入参
	_param *RefundForAikangDoctorRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.doctor.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 退款入参
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) SetParam(_param *RefundForAikangDoctorRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetParam() *RefundForAikangDoctorRequest {
	return r._param
}
