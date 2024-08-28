package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskUpdate 更新轻任务
// taobao.qianniu.task.update
//
// 由任务执行者调用，sub_status，tag和memo至少提供一个
func TaobaoQianniuTaskUpdate(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskUpdateAPIRequest, resp *qianniu.TaobaoQianniuTaskUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
