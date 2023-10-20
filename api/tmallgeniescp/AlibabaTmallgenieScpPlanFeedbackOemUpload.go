package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanfeedbackoemupload 14-供应商反馈（OEM）同步接口
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
func Alibabatmallgeniescpplanfeedbackoemupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanfeedbackoemuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanfeedbackoemuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
