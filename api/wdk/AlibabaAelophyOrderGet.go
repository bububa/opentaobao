package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderGet 翱象拉取订单接口
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
func AlibabaAelophyOrderGet(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderGetAPIRequest, resp *wdk.AlibabaAelophyOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
