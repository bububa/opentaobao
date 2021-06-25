package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销商品数据同步 APIRequest
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest struct {
    model.Params

    // 淘鲜达活动商品信息
    activitySkuList   []DrfTxdActivitySkuBo 

    // 数据版本Id
    versionId   int64 

    // 大润发活动Id
    activityId   string 

    // 活动对应的门店Id
    shopId   string 

}

func NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest{
    return &AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.darunfa.activity.sku.sync"
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) SetActivitySkuList(activitySkuList []DrfTxdActivitySkuBo) error {
    r.activitySkuList = activitySkuList
    r.Set("activity_sku_list", activitySkuList)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetActivitySkuList() []DrfTxdActivitySkuBo {
    return r.activitySkuList
}

func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetVersionId() int64 {
    return r.versionId
}

func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetActivityId() string {
    return r.activityId
}

func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest) GetShopId() string {
    return r.shopId
}

