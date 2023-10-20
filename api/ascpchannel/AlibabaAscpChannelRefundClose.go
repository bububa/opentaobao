package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundClose 渠道退款单关闭
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
func AlibabaAscpChannelRefundClose(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundCloseAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelRefundCloseAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelRefundCloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
