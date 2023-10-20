package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderqueryorder 查询订单
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
func Youkuottpayorderqueryorder(clt *core.SDKClient, req *ottpay.YoukuottpayorderqueryorderAPIRequest, session string) (*ottpay.YoukuottpayorderqueryorderAPIResponse, error) {
	var resp ottpay.YoukuottpayorderqueryorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
