package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnoseordersubmit 体检报告人工解读订单
// alibaba.alihealth.examination.report.diagnose.order.submit
//
// 体检报告人工解读订单信息推送给ISV，进行人工解读
func Alibabaalihealthexaminationreportdiagnoseordersubmit(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnoseordersubmitAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnoseordersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
