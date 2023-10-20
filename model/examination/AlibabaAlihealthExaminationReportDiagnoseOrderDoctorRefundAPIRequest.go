package examination

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.doctor.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundRequest()
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest
func GetAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest 将 AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest(v *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest.Put(v)
}
