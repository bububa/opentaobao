package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityCreate 互动任务活动创建接口
// taobao.jst.interactive.activity.create
//
// 调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
func TaobaoJstInteractiveActivityCreate(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityCreateAPIRequest, resp *jstinteractive.TaobaoJstInteractiveActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
