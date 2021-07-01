package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPatientinfoSyncAPIRequest
处方外流-患者基础信息同步 API请求
alibaba.alihealth.outflow.patientinfo.sync

阿里健康-处方外流-对外提供同步患者基础信息功能 */
type AlibabaAlihealthOutflowPatientinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncPatientInfoRequest *SyncPatientInfoRequest
}

// New
