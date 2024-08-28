package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundClose 渠道退款单关闭
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
func AlibabaAscpChannelRefundClose(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundCloseAPIRequest, resp *ascpchannel.AlibabaAscpChannelRefundCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
