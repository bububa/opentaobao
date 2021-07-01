package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

/* YoukuOttPayOrderQuerycporder
查询支付订单对应cp订单号
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号 */
func YoukuOttPayOrderQuerycporder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQuerycporderAPIRequest, session string) (*ottpay.YoukuOttPayOrderQuerycporderAPIResponse, error) {
	var resp ottpay.YoukuOttPayOrderQuerycporderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
