package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRxCaPrescribeSignedStatusSave 处方ca认证
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
func AlibabaAlihealthRxCaPrescribeSignedStatusSave(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
