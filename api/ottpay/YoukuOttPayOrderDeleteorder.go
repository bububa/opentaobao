package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

/* YoukuOttPayOrderDeleteorder
退订应用中心支付订单
youku.ott.pay.order.deleteorder

应用中心sdk连续包月退订接口 */
func YoukuOttPayOrderDeleteorder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderDeleteorderAPIRequest, session string) (*ottpay.YoukuOttPayOrderDeleteorderAPIResponse, error) {
	var resp ottpay.YoukuOttPayOrderDeleteorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
