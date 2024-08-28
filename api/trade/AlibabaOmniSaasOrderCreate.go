package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaOmniSaasOrderCreate 订单创建接口
// alibaba.omni.saas.order.create
//
// 服务商利用现有的saas系统和阿里完成交易系统的对接
func AlibabaOmniSaasOrderCreate(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaOmniSaasOrderCreateAPIRequest, resp *trade.AlibabaOmniSaasOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
