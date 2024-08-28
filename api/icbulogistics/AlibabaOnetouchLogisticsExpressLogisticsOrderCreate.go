package icbulogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreate 快递下单
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
func AlibabaOnetouchLogisticsExpressLogisticsOrderCreate(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
