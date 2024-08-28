package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOnceTokenGet 网关一次性token获取
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
func TaobaoTopOnceTokenGet(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopOnceTokenGetAPIRequest, resp *tbtrade.TaobaoTopOnceTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
