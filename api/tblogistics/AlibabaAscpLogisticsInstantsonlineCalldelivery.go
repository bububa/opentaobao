package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineCalldelivery 同城配送在线下单正式下单呼叫运力
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
func AlibabaAscpLogisticsInstantsonlineCalldelivery(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
