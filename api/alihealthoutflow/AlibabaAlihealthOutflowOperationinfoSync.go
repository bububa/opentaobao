package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowOperationinfoSync 处方外流-操作信息同步
// alibaba.alihealth.outflow.operationinfo.sync
//
// 阿里健康-处方外流-对外提供同步操作信息功能
func AlibabaAlihealthOutflowOperationinfoSync(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowOperationinfoSyncAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowOperationinfoSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
