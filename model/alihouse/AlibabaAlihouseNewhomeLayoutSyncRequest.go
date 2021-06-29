package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房通户型数据同步 APIRequest
alibaba.alihouse.newhome.layout.sync

房通户型数据同步
*/
type AlibabaAlihouseNewhomeLayoutSyncRequest struct {
    model.Params

    // 数据
    syncProjectLayoutData   *SyncProjectLayoutDto 

}

func NewAlibabaAlihouseNewhomeLayoutSyncRequest() *AlibabaAlihouseNewhomeLayoutSyncRequest{
    return &AlibabaAlihouseNewhomeLayoutSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.layout.sync"
}

func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeLayoutSyncRequest) SetSyncProjectLayoutData(syncProjectLayoutData *SyncProjectLayoutDto) error {
    r.syncProjectLayoutData = syncProjectLayoutData
    r.Set("sync_project_layout_data", syncProjectLayoutData)
    return nil
}

func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetSyncProjectLayoutData() *SyncProjectLayoutDto {
    return r.syncProjectLayoutData
}

