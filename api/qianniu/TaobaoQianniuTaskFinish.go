package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskFinish 完成轻任务
// taobao.qianniu.task.finish
//
// 由任务执行者调用
func TaobaoQianniuTaskFinish(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskFinishAPIRequest, resp *qianniu.TaobaoQianniuTaskFinishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
