package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardDelivery 开始配送工单
// tmall.servicecenter.workcard.delivery
//
// 服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
func TmallServicecenterWorkcardDelivery(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardDeliveryAPIRequest, resp *tmallsc.TmallServicecenterWorkcardDeliveryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
