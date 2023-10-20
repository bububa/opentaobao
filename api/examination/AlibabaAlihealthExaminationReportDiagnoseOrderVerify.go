package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseorderverify 报告解读令牌校验
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
func Alibabaalihealthexaminationreportdiagnoseorderverify(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseorderverifyAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseorderverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
