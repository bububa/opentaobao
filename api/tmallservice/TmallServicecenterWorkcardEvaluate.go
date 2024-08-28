package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardEvaluate 服务商反馈鉴定结果
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
func TmallServicecenterWorkcardEvaluate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardEvaluateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardEvaluateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
