package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
数据同步版本号申请 API请求
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
type AlibabaWdkMarketingOpenVersionApplyRequest struct {
    model.Params
    // 同步版本信息
    syncVersion   *SyncVersionBo
}

// 初始化AlibabaWdkMarketingOpenVersionApplyRequest对象
func NewAlibabaWdkMarketingOpenVersionApplyRequest() *AlibabaWdkMarketingOpenVersionApplyRequest{
    return &AlibabaWdkMarketingOpenVersionApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.version.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncVersion Setter
// 同步版本信息
func (r *AlibabaWdkMarketingOpenVersionApplyRequest) SetSyncVersion(syncVersion *SyncVersionBo) error {
    r.syncVersion = syncVersion
    r.Set("sync_version", syncVersion)
    return nil
}

// SyncVersion Getter
func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetSyncVersion() *SyncVersionBo {
    return r.syncVersion
}
