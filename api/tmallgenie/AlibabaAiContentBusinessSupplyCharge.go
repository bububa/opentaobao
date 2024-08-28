package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSupplyCharge 供销商品充值接口
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
func AlibabaAiContentBusinessSupplyCharge(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSupplyChargeAPIRequest, resp *tmallgenie.AlibabaAiContentBusinessSupplyChargeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
