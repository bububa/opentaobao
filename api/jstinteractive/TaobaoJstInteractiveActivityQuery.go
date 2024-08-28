package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityQuery 互动任务活动查询接口
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
func TaobaoJstInteractiveActivityQuery(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractiveActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
