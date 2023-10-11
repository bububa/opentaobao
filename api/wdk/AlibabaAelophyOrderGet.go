package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderGet 翱象拉取订单接口
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
func AlibabaAelophyOrderGet(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderGetAPIRequest, session string) (*wdk.AlibabaAelophyOrderGetAPIResponse, error) {
	var resp wdk.AlibabaAelophyOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
