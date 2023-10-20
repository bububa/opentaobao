package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderDelivererChange 配送员信息变更接口
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
func AlibabaAelophyOrderDelivererChange(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderDelivererChangeAPIRequest, resp *wdk.AlibabaAelophyOrderDelivererChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
