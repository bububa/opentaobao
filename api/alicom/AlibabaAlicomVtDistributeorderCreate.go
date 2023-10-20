package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeorderCreate 通信业务外放下单
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
func AlibabaAlicomVtDistributeorderCreate(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeorderCreateAPIRequest, resp *alicom.AlibabaAlicomVtDistributeorderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
