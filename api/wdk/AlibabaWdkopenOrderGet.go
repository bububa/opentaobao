package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkopenorderget 五道口商户订单获取
// alibaba.wdkopen.order.get
//
// 商户通过五道口订单id获取订单信息
func Alibabawdkopenorderget(clt *core.SDKClient, req *wdk.AlibabawdkopenordergetAPIRequest, session string) (*wdk.AlibabawdkopenordergetAPIResponse, error) {
	var resp wdk.AlibabawdkopenordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
