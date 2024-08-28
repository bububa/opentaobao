package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveTaskQuery 互动任务列表查询接口
// taobao.jst.interactive.task.query
//
// 查询用户的互动任务列表
func TaobaoJstInteractiveTaskQuery(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveTaskQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractiveTaskQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
