package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseordersummary 体检报告人工解读总结回传
// alibaba.alihealth.examination.report.diagnose.order.summary
//
// 记录体检报告人工解读总结
func Alibabaalihealthexaminationreportdiagnoseordersummary(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseordersummaryAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseordersummaryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
