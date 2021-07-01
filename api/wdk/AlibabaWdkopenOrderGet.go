package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkopenOrderGet
五道口商户订单获取
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息 */
func AlibabaWdkopenOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkopenOrderGetAPIRequest, session string) (*wdk.AlibabaWdkopenOrderGetAPIResponse, error) {
	var resp wdk.AlibabaWdkopenOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
