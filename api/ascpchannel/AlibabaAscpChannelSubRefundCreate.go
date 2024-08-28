package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSubRefundCreate 淘外分销逆向创单（子单退）
// alibaba.ascp.channel.sub.refund.create
//
// 淘外分销逆向创单（子单退）
func AlibabaAscpChannelSubRefundCreate(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSubRefundCreateAPIRequest, resp *ascpchannel.AlibabaAscpChannelSubRefundCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
