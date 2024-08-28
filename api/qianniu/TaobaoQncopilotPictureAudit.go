package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQncopilotPictureAudit AIGC创作图片审核
// taobao.qncopilot.picture.audit
//
// AIGC创作图片审核
func TaobaoQncopilotPictureAudit(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQncopilotPictureAuditAPIRequest, resp *qianniu.TaobaoQncopilotPictureAuditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
