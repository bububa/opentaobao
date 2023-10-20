package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeQueryprotocol 通信业务外放协议查询
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
func AlibabaAlicomVtDistributeQueryprotocol(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeQueryprotocolAPIRequest, resp *alicom.AlibabaAlicomVtDistributeQueryprotocolAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
