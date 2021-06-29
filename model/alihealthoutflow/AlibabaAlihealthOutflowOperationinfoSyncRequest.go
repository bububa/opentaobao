package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-操作信息同步 APIRequest
alibaba.alihealth.outflow.operationinfo.sync

阿里健康-处方外流-对外提供同步操作信息功能
*/
type AlibabaAlihealthOutflowOperationinfoSyncRequest struct {
    model.Params

    // 入参
    syncOperationInfoRequest   *SyncOperationInfoRequest 

}

func NewAlibabaAlihealthOutflowOperationinfoSyncRequest() *AlibabaAlihealthOutflowOperationinfoSyncRequest{
    return &AlibabaAlihealthOutflowOperationinfoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowOperationinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.operationinfo.sync"
}

func (r AlibabaAlihealthOutflowOperationinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowOperationinfoSyncRequest) SetSyncOperationInfoRequest(syncOperationInfoRequest *SyncOperationInfoRequest) error {
    r.syncOperationInfoRequest = syncOperationInfoRequest
    r.Set("sync_operation_info_request", syncOperationInfoRequest)
    return nil
}

func (r AlibabaAlihealthOutflowOperationinfoSyncRequest) GetSyncOperationInfoRequest() *SyncOperationInfoRequest {
    return r.syncOperationInfoRequest
}

