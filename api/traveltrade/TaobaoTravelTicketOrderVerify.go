package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoTravelTicketOrderVerify 飞猪门票核销通知
// taobao.travel.ticket.order.verify
//
// 系统商通过TOP接口调用通知飞猪门票核销情况
func TaobaoTravelTicketOrderVerify(ctx context.Context, clt *core.SDKClient, req *traveltrade.TaobaoTravelTicketOrderVerifyAPIRequest, resp *traveltrade.TaobaoTravelTicketOrderVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
