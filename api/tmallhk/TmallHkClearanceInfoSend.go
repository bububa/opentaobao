package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceInfoSend 清关信息回调通知
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
func TmallHkClearanceInfoSend(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallHkClearanceInfoSendAPIRequest, resp *tmallhk.TmallHkClearanceInfoSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
