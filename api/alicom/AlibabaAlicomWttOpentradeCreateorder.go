package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomWttOpentradeCreateorder 充值送活动下单接口
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
func AlibabaAlicomWttOpentradeCreateorder(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomWttOpentradeCreateorderAPIRequest, resp *alicom.AlibabaAlicomWttOpentradeCreateorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
