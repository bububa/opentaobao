package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelhotelticketorderrefund 退款结结果通知
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
func Alitriptravelhotelticketorderrefund(clt *core.SDKClient, req *traveltrade.AlitriptravelhotelticketorderrefundAPIRequest, session string) (*traveltrade.AlitriptravelhotelticketorderrefundAPIResponse, error) {
	var resp traveltrade.AlitriptravelhotelticketorderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
