package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowVisitinfoSyncAPIRequest
处方外流-问诊、处方同步 API请求
alibaba.alihealth.outflow.visitinfo.sync

阿里健康-处方外流-对外提供问诊、处方功能 */
type AlibabaAlihealthOutflowVisitinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncVisitInfoRequest *SyncVisitInfoRequest
}

// New
