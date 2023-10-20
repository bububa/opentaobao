package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderCreate 创建订单
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
func AlibabaWdkChannelOrderCreate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderCreateAPIRequest, resp *wdk.AlibabaWdkChannelOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
