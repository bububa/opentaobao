package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanfeedbackrawupload 15-供应商反馈（原料）同步接口
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
func Alibabatmallgeniescpplanfeedbackrawupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanfeedbackrawuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanfeedbackrawuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
