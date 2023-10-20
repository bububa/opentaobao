package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefund 报告解读订单-医生退款
// alibaba.alihealth.examination.report.diagnose.order.doctor.refund
//
// 报告解读订单医生退款
func AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefund(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
