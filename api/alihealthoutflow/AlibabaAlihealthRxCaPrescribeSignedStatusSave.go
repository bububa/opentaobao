package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRxCaPrescribeSignedStatusSave 处方ca认证
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
func AlibabaAlihealthRxCaPrescribeSignedStatusSave(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
