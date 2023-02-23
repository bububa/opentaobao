package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPatientinfoSyncAPIRequest 处方外流-患者基础信息同步 API请求
// alibaba.alihealth.outflow.patientinfo.sync
//
// 阿里健康-处方外流-对外提供同步患者基础信息功能
type AlibabaAlihealthOutflowPatientinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncPatientInfoRequest *SyncPatientInfoRequest
}

// NewAlibabaAlihealthOutflowPatientinfoSyncRequest 初始化AlibabaAlihealthOutflowPatientinfoSyncAPIRequest对象
func NewAlibabaAlihealthOutflowPatientinfoSyncRequest() *AlibabaAlihealthOutflowPatientinfoSyncAPIRequest {
	return &AlibabaAlihealthOutflowPatientinfoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.patientinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncPatientInfoRequest is SyncPatientInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) SetSyncPatientInfoRequest(_syncPatientInfoRequest *SyncPatientInfoRequest) error {
	r._syncPatientInfoRequest = _syncPatientInfoRequest
	r.Set("sync_patient_info_request", _syncPatientInfoRequest)
	return nil
}

// GetSyncPatientInfoRequest SyncPatientInfoRequest Getter
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetSyncPatientInfoRequest() *SyncPatientInfoRequest {
	return r._syncPatientInfoRequest
}
