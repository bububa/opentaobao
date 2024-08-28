package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTasksCount 任务查询条数接口
// taobao.qianniu.tasks.count
//
// 任务查询条数接口
func TaobaoQianniuTasksCount(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTasksCountAPIRequest, resp *qianniu.TaobaoQianniuTasksCountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
