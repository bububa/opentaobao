package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundGoodsWaybill 淘外分销退货回传物流单号
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
func AlibabaAscpChannelRefundGoodsWaybill(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundGoodsWaybillAPIRequest, resp *ascpchannel.AlibabaAscpChannelRefundGoodsWaybillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
