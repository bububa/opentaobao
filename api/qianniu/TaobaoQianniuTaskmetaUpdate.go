package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskmetaUpdate 更新任务元数据
// taobao.qianniu.taskmeta.update
//
// 由任务发起者调用
func TaobaoQianniuTaskmetaUpdate(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetaUpdateAPIRequest, resp *qianniu.TaobaoQianniuTaskmetaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
