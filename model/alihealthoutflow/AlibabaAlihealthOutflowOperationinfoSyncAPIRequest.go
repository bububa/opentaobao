package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowOperationinfoSyncAPIRequest
处方外流-操作信息同步 API请求
alibaba.alihealth.outflow.operationinfo.sync

阿里健康-处方外流-对外提供同步操作信息功能 */
type AlibabaAlihealthOutflowOperationinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncOperationInfoRequest *SyncOperationInfoRequest
}

// New
