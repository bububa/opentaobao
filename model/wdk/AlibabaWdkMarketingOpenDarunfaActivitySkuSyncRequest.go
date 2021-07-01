package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销商品数据同步 API请求
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest struct {
    model.Params
    // 淘鲜达活动商品信息
    _activitySkuList   []DrfTxdActivitySkuBo
    // 数据版本Id
    _versionId   int64
    // 大润发活动Id
    _activityId   string
    // 活动对应的门店Id
    _shopId   string
}

// 初始化AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest对象
func NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest{
    return &AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.darunfa.activity.sku.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivitySkuList Setter
// 淘鲜达活动商品信息
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivitySkuList(_activitySkuList []DrfTxdActivitySkuBo) error {
    r._activitySkuList = _activitySkuList
    r.Set("activity_sku_list", _activitySkuList)
    return nil
}

// ActivitySkuList Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivitySkuList() []DrfTxdActivitySkuBo {
    return r._activitySkuList
}
// VersionId Setter
// 数据版本Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetVersionId() int64 {
    return r._versionId
}
// ActivityId Setter
// 大润发活动Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivityId() string {
    return r._activityId
}
// ShopId Setter
// 活动对应的门店Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetShopId() string {
    return r._shopId
}
