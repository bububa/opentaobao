package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseorderdoctorrefund 报告解读订单-医生退款
// alibaba.alihealth.examination.report.diagnose.order.doctor.refund
//
// 报告解读订单医生退款
func Alibabaalihealthexaminationreportdiagnoseorderdoctorrefund(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseorderdoctorrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
