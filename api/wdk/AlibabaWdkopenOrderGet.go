package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkopenOrderGet 五道口商户订单获取
// alibaba.wdkopen.order.get
//
// 商户通过五道口订单id获取订单信息
func AlibabaWdkopenOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkopenOrderGetAPIRequest, resp *wdk.AlibabaWdkopenOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
