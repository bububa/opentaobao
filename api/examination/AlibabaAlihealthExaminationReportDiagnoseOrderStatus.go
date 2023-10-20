package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseorderstatus 报告解读订单状态更新
// alibaba.alihealth.examination.report.diagnose.order.status
//
// 报告解读订单状态更新
func Alibabaalihealthexaminationreportdiagnoseorderstatus(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseorderstatusAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseorderstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
