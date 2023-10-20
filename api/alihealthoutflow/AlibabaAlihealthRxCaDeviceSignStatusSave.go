package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRxCaDeviceSignStatusSave 处方ca认证-厂商通知接口
// alibaba.alihealth.rx.ca.device.sign.status.save
//
// 处方ca认证-厂商通知接口
func AlibabaAlihealthRxCaDeviceSignStatusSave(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRxCaDeviceSignStatusSaveAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
