package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionCreate 处方外流-创建处方
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
func AlibabaAlihealthOutflowPrescriptionCreate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionCreateAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
