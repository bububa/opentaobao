package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreportdiagnosetempmessagereceive 导医通报告解读临时消息接收
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
func Alibabaalihealthexaminationreportdiagnosetempmessagereceive(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest, session string) (*examination.AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
