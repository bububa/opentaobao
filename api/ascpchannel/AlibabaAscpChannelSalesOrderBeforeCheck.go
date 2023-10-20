package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsalesorderbeforecheck 供应链外部订单创建前校验接口
// alibaba.ascp.channel.sales.order.before.check
//
// 供应链外部订单创建前校验接口
func Alibabaascpchannelsalesorderbeforecheck(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsalesorderbeforecheckAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsalesorderbeforecheckAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsalesorderbeforecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
