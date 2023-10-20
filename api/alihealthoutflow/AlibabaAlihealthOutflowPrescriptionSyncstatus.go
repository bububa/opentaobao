package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionSyncstatus 处方外流-处方状态同步
// alibaba.alihealth.outflow.prescription.syncstatus
//
// 阿里健康-处方外流-对外提供同步处方状态功能
func AlibabaAlihealthOutflowPrescriptionSyncstatus(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionSyncstatusAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
