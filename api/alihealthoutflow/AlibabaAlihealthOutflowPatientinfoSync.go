package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPatientinfoSync 处方外流-患者基础信息同步
// alibaba.alihealth.outflow.patientinfo.sync
//
// 阿里健康-处方外流-对外提供同步患者基础信息功能
func AlibabaAlihealthOutflowPatientinfoSync(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPatientinfoSyncAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPatientinfoSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
