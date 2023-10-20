package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundCancel 渠道退款单撤销
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
func AlibabaAscpChannelRefundCancel(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundCancelAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelRefundCancelAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelRefundCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
