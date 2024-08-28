package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskmetasGet 任务元查询接口
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
func TaobaoQianniuTaskmetasGet(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetasGetAPIRequest, resp *qianniu.TaobaoQianniuTaskmetasGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
