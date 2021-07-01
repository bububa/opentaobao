package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

/* YoukuOttPayOrderCreateorder
创建订单
youku.ott.pay.order.createorder

ottpay创建订单 */
func YoukuOttPayOrderCreateorder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderCreateorderAPIRequest, session string) (*ottpay.YoukuOttPayOrderCreateorderAPIResponse, error) {
	var resp ottpay.YoukuOttPayOrderCreateorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
