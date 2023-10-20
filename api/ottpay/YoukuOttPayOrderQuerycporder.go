package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderquerycporder 查询支付订单对应cp订单号
// youku.ott.pay.order.querycporder
//
// 根据支付订单查询对应cp订单号
func Youkuottpayorderquerycporder(clt *core.SDKClient, req *ottpay.YoukuottpayorderquerycporderAPIRequest, session string) (*ottpay.YoukuottpayorderquerycporderAPIResponse, error) {
	var resp ottpay.YoukuottpayorderquerycporderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
