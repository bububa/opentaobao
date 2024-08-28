package refund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundsReceiveGet 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func TaobaoRefundsReceiveGet(ctx context.Context, clt *core.SDKClient, req *refund.TaobaoRefundsReceiveGetAPIRequest, resp *refund.TaobaoRefundsReceiveGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
