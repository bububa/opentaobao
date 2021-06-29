package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-患者基础信息同步 API请求
alibaba.alihealth.outflow.patientinfo.sync

阿里健康-处方外流-对外提供同步患者基础信息功能
*/
type AlibabaAlihealthOutflowPatientinfoSyncRequest struct {
    model.Params
    // 入参
    syncPatientInfoRequest   *SyncPatientInfoRequest
}

// 初始化AlibabaAlihealthOutflowPatientinfoSyncRequest对象
func NewAlibabaAlihealthOutflowPatientinfoSyncRequest() *AlibabaAlihealthOutflowPatientinfoSyncRequest{
    return &AlibabaAlihealthOutflowPatientinfoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.patientinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncPatientInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowPatientinfoSyncRequest) SetSyncPatientInfoRequest(syncPatientInfoRequest *SyncPatientInfoRequest) error {
    r.syncPatientInfoRequest = syncPatientInfoRequest
    r.Set("sync_patient_info_request", syncPatientInfoRequest)
    return nil
}

// SyncPatientInfoRequest Getter
func (r AlibabaAlihealthOutflowPatientinfoSyncRequest) GetSyncPatientInfoRequest() *SyncPatientInfoRequest {
    return r.syncPatientInfoRequest
}
