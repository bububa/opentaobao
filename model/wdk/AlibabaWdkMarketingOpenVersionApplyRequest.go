package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
数据同步版本号申请 APIRequest
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
type AlibabaWdkMarketingOpenVersionApplyRequest struct {
    model.Params

    // 同步版本信息
    syncVersion   *SyncVersionBo 

}

func NewAlibabaWdkMarketingOpenVersionApplyRequest() *AlibabaWdkMarketingOpenVersionApplyRequest{
    return &AlibabaWdkMarketingOpenVersionApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.version.apply"
}

func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenVersionApplyRequest) SetSyncVersion(syncVersion *SyncVersionBo) error {
    r.syncVersion = syncVersion
    r.Set("sync_version", syncVersion)
    return nil
}

func (r AlibabaWdkMarketingOpenVersionApplyRequest) GetSyncVersion() *SyncVersionBo {
    return r.syncVersion
}

