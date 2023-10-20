package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeQueryprotocol 通信业务外放协议查询
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
func AlibabaAlicomVtDistributeQueryprotocol(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeQueryprotocolAPIRequest, session string) (*alicom.AlibabaAlicomVtDistributeQueryprotocolAPIResponse, error) {
	var resp alicom.AlibabaAlicomVtDistributeQueryprotocolAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
