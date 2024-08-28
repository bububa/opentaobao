package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsOfflineSend 自己联系物流发货
// alibaba.ascp.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
func AlibabaAscpLogisticsOfflineSend(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsOfflineSendAPIRequest, resp *tblogistics.AlibabaAscpLogisticsOfflineSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
