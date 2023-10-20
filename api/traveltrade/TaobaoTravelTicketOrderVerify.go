package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoTravelTicketOrderVerify 飞猪门票核销通知
// taobao.travel.ticket.order.verify
//
// 系统商通过TOP接口调用通知飞猪门票核销情况
func TaobaoTravelTicketOrderVerify(clt *core.SDKClient, req *traveltrade.TaobaoTravelTicketOrderVerifyAPIRequest, session string) (*traveltrade.TaobaoTravelTicketOrderVerifyAPIResponse, error) {
	var resp traveltrade.TaobaoTravelTicketOrderVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
