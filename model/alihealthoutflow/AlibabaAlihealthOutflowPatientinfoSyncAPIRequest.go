package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowpatientinfosyncAPIRequest 处方外流-患者基础信息同步 API请求
// alibaba.alihealth.outflow.patientinfo.sync
//
// 阿里健康-处方外流-对外提供同步患者基础信息功能
type AlibabaalihealthoutflowpatientinfosyncAPIRequest struct {
	model.Params
	// 入参
	_syncPatientInfoRequest *SyncPatientInfoRequest
}

// NewAlibabaalihealthoutflowpatientinfosyncRequest 初始化AlibabaalihealthoutflowpatientinfosyncAPIRequest对象
func NewAlibabaalihealthoutflowpatientinfosyncRequest() *AlibabaalihealthoutflowpatientinfosyncAPIRequest {
	return &AlibabaalihealthoutflowpatientinfosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowpatientinfosyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.patientinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowpatientinfosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowpatientinfosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPatientInfoRequest is SyncPatientInfoRequest Setter
// 入参
func (r *AlibabaalihealthoutflowpatientinfosyncAPIRequest) SetSyncPatientInfoRequest(_syncPatientInfoRequest *SyncPatientInfoRequest) error {
	r._syncPatientInfoRequest = _syncPatientInfoRequest
	r.Set("sync_patient_info_request", _syncPatientInfoRequest)
	return nil
}

// GetSyncPatientInfoRequest SyncPatientInfoRequest Getter
func (r AlibabaalihealthoutflowpatientinfosyncAPIRequest) GetSyncPatientInfoRequest() *SyncPatientInfoRequest {
	return r._syncPatientInfoRequest
}
