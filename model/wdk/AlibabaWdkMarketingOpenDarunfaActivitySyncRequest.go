package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动数据同步 APIRequest
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySyncRequest struct {
    model.Params

    // 大润发活动数据
    activityList   []DrfTxdActivityBo 

    // 门店Id
    shopId   string 

    // 版本ID
    versionId   int64 

}

func NewAlibabaWdkMarketingOpenDarunfaActivitySyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySyncRequest{
    return &AlibabaWdkMarketingOpenDarunfaActivitySyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.darunfa.activity.sync"
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) SetActivityList(activityList []DrfTxdActivityBo) error {
    r.activityList = activityList
    r.Set("activity_list", activityList)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) GetActivityList() []DrfTxdActivityBo {
    return r.activityList
}

func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) GetShopId() string {
    return r.shopId
}

func (r *AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySyncRequest) GetVersionId() int64 {
    return r.versionId
}

