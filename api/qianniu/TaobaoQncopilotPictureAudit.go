package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQncopilotPictureAudit AIGC创作图片审核
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
func TaobaoQncopilotPictureAudit(clt *core.SDKClient, req *qianniu.TaobaoQncopilotPictureAuditAPIRequest, session string) (*qianniu.TaobaoQncopilotPictureAuditAPIResponse, error) {
	var resp qianniu.TaobaoQncopilotPictureAuditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
