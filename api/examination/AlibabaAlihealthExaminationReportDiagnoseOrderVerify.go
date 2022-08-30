package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderVerify 报告解读令牌校验
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
func AlibabaAlihealthExaminationReportDiagnoseOrderVerify(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
