package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

/* TaobaoTravelTicketOrderRefund
飞猪门票退票结果通知
taobao.travel.ticket.order.refund

门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。 */
func TaobaoTravelTicketOrderRefund(clt *core.SDKClient, req *traveltrade.TaobaoTravelTicketOrderRefundAPIRequest, session string) (*traveltrade.TaobaoTravelTicketOrderRefundAPIResponse, error) {
	var resp traveltrade.TaobaoTravelTicketOrderRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
