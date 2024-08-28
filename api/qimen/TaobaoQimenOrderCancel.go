package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderCancel 单据取消接口
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
func TaobaoQimenOrderCancel(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderCancelAPIRequest, resp *qimen.TaobaoQimenOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
