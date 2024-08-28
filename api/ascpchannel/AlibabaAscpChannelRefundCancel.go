package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundCancel 渠道退款单撤销
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
func AlibabaAscpChannelRefundCancel(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundCancelAPIRequest, resp *ascpchannel.AlibabaAscpChannelRefundCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
