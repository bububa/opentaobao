package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelMainRefundCreate 淘外分销逆向创单（未发货整单退）
// alibaba.ascp.channel.main.refund.create
//
// 淘外分销解决方案--订单--逆向创单（未发货整单退）
func AlibabaAscpChannelMainRefundCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelMainRefundCreateAPIRequest, resp *ascpchannel.AlibabaAscpChannelMainRefundCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
