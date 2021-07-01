package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPrescriptionSyncstatusAPIRequest
处方外流-处方状态同步 API请求
alibaba.alihealth.outflow.prescription.syncstatus

阿里健康-处方外流-对外提供同步处方状态功能 */
type AlibabaAlihealthOutflowPrescriptionSyncstatusAPIRequest struct {
	model.Params
	// 入参
	_syncStatusRequest *SyncPrescriptionStatusRequest
}

// New
