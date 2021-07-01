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
type AlibabaAlihouseNewhomeLayoutSyncAPIRequest struct {
    model.Params
    // 数据
    _syncProjectLayoutData   *SyncProjectLayoutDTO
}

// 初始化AlibabaAlihouseNewhomeLayoutSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeLayoutSyncRequest() *AlibabaAlihouseNewhomeLayoutSyncAPIRequest{
    return &AlibabaAlihouseNewhomeLayoutSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.layout.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncProjectLayoutData Setter
// 数据
func (r *AlibabaAlihouseNewhomeLayoutSyncAPIRequest) SetSyncProjectLayoutData(_syncProjectLayoutData *SyncProjectLayoutDTO) error {
    r._syncProjectLayoutData = _syncProjectLayoutData
    r.Set("sync_project_layout_data", _syncProjectLayoutData)
    return nil
}

// SyncProjectLayoutData Getter
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetSyncProjectLayoutData() *SyncProjectLayoutDTO {
    return r._syncProjectLayoutData
}
