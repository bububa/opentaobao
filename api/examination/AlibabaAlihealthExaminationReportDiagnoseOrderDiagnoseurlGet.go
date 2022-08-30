package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGet 获取报告解读url
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
func AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGet(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
