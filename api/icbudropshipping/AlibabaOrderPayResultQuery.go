package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabaorderpayresultquery alibaba查询订单支付结果
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
func Alibabaorderpayresultquery(clt *core.SDKClient, req *icbudropshipping.AlibabaorderpayresultqueryAPIRequest, session string) (*icbudropshipping.AlibabaorderpayresultqueryAPIResponse, error) {
	var resp icbudropshipping.AlibabaorderpayresultqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
