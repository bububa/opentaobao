package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeorderCreate 通信业务外放下单
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
func AlibabaAlicomVtDistributeorderCreate(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeorderCreateAPIRequest, resp *alicom.AlibabaAlicomVtDistributeorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
