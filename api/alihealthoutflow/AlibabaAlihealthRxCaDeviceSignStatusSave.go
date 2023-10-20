package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRxCaDeviceSignStatusSave 处方ca认证-厂商通知接口
// alibaba.alihealth.rx.ca.device.sign.status.save
//
// 处方ca认证-厂商通知接口
func AlibabaAlihealthRxCaDeviceSignStatusSave(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
