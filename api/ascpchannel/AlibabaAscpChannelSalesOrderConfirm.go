package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsalesorderconfirm 渠道销售单确认收货
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
func Alibabaascpchannelsalesorderconfirm(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsalesorderconfirmAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsalesorderconfirmAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsalesorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
