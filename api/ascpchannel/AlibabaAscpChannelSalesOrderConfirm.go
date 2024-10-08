package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSalesOrderConfirm 渠道销售单确认收货
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
func AlibabaAscpChannelSalesOrderConfirm(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSalesOrderConfirmAPIRequest, resp *ascpchannel.AlibabaAscpChannelSalesOrderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
