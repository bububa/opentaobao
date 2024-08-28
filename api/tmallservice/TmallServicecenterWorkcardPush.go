package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardPush 推送服务工单信息
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
func TmallServicecenterWorkcardPush(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardPushAPIRequest, resp *tmallservice.TmallServicecenterWorkcardPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
