package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGet 同城配送在线下单获取配送单
// alibaba.ascp.logistics.instantsonline.deliveryorder.get
//
// 同城配送在线下单获取配送单
func AlibabaAscpLogisticsInstantsonlineDeliveryorderGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
