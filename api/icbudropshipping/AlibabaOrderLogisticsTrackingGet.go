package icbudropshipping

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaOrderLogisticsTrackingGet 阿里巴巴订单物流轨迹查询
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
func AlibabaOrderLogisticsTrackingGet(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIRequest, resp *icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
