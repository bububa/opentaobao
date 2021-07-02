package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSubRefundCreate 淘外分销逆向创单（子单退）
// alibaba.ascp.channel.sub.refund.create
//
// 淘外分销逆向创单（子单退）
func AlibabaAscpChannelSubRefundCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSubRefundCreateAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSubRefundCreateAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSubRefundCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
