package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelhotelticketordercreate 创单(支付订单)通知
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
func Alitriptravelhotelticketordercreate(clt *core.SDKClient, req *traveltrade.AlitriptravelhotelticketordercreateAPIRequest, session string) (*traveltrade.AlitriptravelhotelticketordercreateAPIResponse, error) {
	var resp traveltrade.AlitriptravelhotelticketordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
