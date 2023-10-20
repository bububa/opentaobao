package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Alitripiebuyerorderbookpay 【国际机票】下单预定支付
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
func Alitripiebuyerorderbookpay(clt *core.SDKClient, req *ieagency.AlitripiebuyerorderbookpayAPIRequest, session string) (*ieagency.AlitripiebuyerorderbookpayAPIResponse, error) {
	var resp ieagency.AlitripiebuyerorderbookpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
