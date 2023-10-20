package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqncopilotpictureaudit AIGC创作图片审核
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
func Taobaoqncopilotpictureaudit(clt *core.SDKClient, req *qianniu.TaobaoqncopilotpictureauditAPIRequest, session string) (*qianniu.TaobaoqncopilotpictureauditAPIResponse, error) {
	var resp qianniu.TaobaoqncopilotpictureauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
