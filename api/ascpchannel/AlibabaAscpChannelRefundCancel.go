package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelrefundcancel 渠道退款单撤销
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
func Alibabaascpchannelrefundcancel(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelrefundcancelAPIRequest, session string) (*ascpchannel.AlibabaascpchannelrefundcancelAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelrefundcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
