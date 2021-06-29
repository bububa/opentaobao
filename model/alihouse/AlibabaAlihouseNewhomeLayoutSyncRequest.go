package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房通户型数据同步 API请求
alibaba.alihouse.newhome.layout.sync

房通户型数据同步
*/
type AlibabaAlihouseNewhomeLayoutSyncRequest struct {
    model.Params
    // 数据
    _syncProjectLayoutData   *SyncProjectLayoutDTO
}

// 初始化AlibabaAlihouseNewhomeLayoutSyncRequest对象
func NewAlibabaAlihouseNewhomeLayoutSyncRequest() *AlibabaAlihouseNewhomeLayoutSyncRequest{
    return &AlibabaAlihouseNewhomeLayoutSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.layout.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncProjectLayoutData Setter
// 数据
func (r *AlibabaAlihouseNewhomeLayoutSyncRequest) SetSyncProjectLayoutData(_syncProjectLayoutData *SyncProjectLayoutDTO) error {
    r._syncProjectLayoutData = _syncProjectLayoutData
    r.Set("sync_project_layout_data", _syncProjectLayoutData)
    return nil
}

// SyncProjectLayoutData Getter
func (r AlibabaAlihouseNewhomeLayoutSyncRequest) GetSyncProjectLayoutData() *SyncProjectLayoutDTO {
    return r._syncProjectLayoutData
}
