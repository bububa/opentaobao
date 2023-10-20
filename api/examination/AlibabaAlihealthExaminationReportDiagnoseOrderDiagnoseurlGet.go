package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseorderdiagnoseurlget 获取报告解读url
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
func Alibabaalihealthexaminationreportdiagnoseorderdiagnoseurlget(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
