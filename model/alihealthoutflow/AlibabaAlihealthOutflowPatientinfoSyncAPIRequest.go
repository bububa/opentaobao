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
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SyncPatientInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) SetSyncPatientInfoRequest(_syncPatientInfoRequest *SyncPatientInfoRequest) error {
	r._syncPatientInfoRequest = _syncPatientInfoRequest
	r.Set("sync_patient_info_request", _syncPatientInfoRequest)
	return nil
}

// Get SyncPatientInfoRequest Getter
func (r AlibabaAlihealthOutflowPatientinfoSyncAPIRequest) GetSyncPatientInfoRequest() *SyncPatientInfoRequest {
	return r._syncPatientInfoRequest
}
