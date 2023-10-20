package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenorderbackflow 订单回流接口
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
func Alibabaalsccrmopenorderbackflow(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenorderbackflowAPIRequest, session string) (*alsc.AlibabaalsccrmopenorderbackflowAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenorderbackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
