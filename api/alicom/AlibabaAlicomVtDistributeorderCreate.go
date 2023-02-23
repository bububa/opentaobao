package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeorderCreate 通信业务外放下单
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
func AlibabaAlicomVtDistributeorderCreate(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeorderCreateAPIRequest, session string) (*alicom.AlibabaAlicomVtDistributeorderCreateAPIResponse, error) {
	var resp alicom.AlibabaAlicomVtDistributeorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
