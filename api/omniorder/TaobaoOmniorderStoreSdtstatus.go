package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSdtstatus 菜鸟裹裹运单状态查询
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
func TaobaoOmniorderStoreSdtstatus(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtstatusAPIRequest, resp *omniorder.TaobaoOmniorderStoreSdtstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
