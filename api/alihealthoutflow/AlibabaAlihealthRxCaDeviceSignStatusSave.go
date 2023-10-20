package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthrxcadevicesignstatussave 处方ca认证-厂商通知接口
// alibaba.alihealth.rx.ca.device.sign.status.save
//
// 处方ca认证-厂商通知接口
func Alibabaalihealthrxcadevicesignstatussave(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthrxcadevicesignstatussaveAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthrxcadevicesignstatussaveAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthrxcadevicesignstatussaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
