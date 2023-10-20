package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowprescriptionsyncstatus 处方外流-处方状态同步
// alibaba.alihealth.outflow.prescription.syncstatus
//
// 阿里健康-处方外流-对外提供同步处方状态功能
func Alibabaalihealthoutflowprescriptionsyncstatus(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowprescriptionsyncstatusAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowprescriptionsyncstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
