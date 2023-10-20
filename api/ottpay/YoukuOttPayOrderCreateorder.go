package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayordercreateorder 创建订单
// youku.ott.pay.order.createorder
//
// ottpay创建订单
func Youkuottpayordercreateorder(clt *core.SDKClient, req *ottpay.YoukuottpayordercreateorderAPIRequest, session string) (*ottpay.YoukuottpayordercreateorderAPIResponse, error) {
	var resp ottpay.YoukuottpayordercreateorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
