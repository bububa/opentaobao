package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderqueryorderbycp 订单查询接口(cp订单号查询)
// youku.ott.pay.order.queryorderbycp
//
// 给商户服务端查询订单状态
func Youkuottpayorderqueryorderbycp(clt *core.SDKClient, req *ottpay.YoukuottpayorderqueryorderbycpAPIRequest, session string) (*ottpay.YoukuottpayorderqueryorderbycpAPIResponse, error) {
	var resp ottpay.YoukuottpayorderqueryorderbycpAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
