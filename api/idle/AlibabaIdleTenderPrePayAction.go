package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderPrePayAction 服务商预付款完成接口
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
func AlibabaIdleTenderPrePayAction(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderPrePayActionAPIRequest, resp *idle.AlibabaIdleTenderPrePayActionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
