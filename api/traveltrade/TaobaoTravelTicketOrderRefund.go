package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoTravelTicketOrderRefund 飞猪门票退票结果通知
// taobao.travel.ticket.order.refund
//
// 门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。
func TaobaoTravelTicketOrderRefund(ctx context.Context, clt *core.SDKClient, req *traveltrade.TaobaoTravelTicketOrderRefundAPIRequest, resp *traveltrade.TaobaoTravelTicketOrderRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
