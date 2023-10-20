package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelhotelticketorderverify 订单核销通知
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
func Alitriptravelhotelticketorderverify(clt *core.SDKClient, req *traveltrade.AlitriptravelhotelticketorderverifyAPIRequest, session string) (*traveltrade.AlitriptravelhotelticketorderverifyAPIResponse, error) {
	var resp traveltrade.AlitriptravelhotelticketorderverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
