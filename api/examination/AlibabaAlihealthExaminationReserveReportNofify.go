package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreservereportnofify 服务商主动通知体检报告
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
func Alibabaalihealthexaminationreservereportnofify(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreservereportnofifyAPIRequest, session string) (*examination.AlibabaalihealthexaminationreservereportnofifyAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreservereportnofifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
