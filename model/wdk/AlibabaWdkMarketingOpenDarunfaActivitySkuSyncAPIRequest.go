package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest 营销商品数据同步 API请求
// alibaba.wdk.marketing.open.darunfa.activity.sku.sync
//
// 大润发营销商品数据同步
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest struct {
	model.Params
	// 淘鲜达活动商品信息
	_activitySkuList []DrfTxdActivitySkuBo
	// 数据版本Id
	_versionId int64
	// 大润发活动Id
	_activityId string
	// 活动对应的门店Id
	_shopId string
}

// NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest 初始化AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest对象
func NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest {
	return &AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.darunfa.activity.sku.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivitySkuList Setter
// 淘鲜达活动商品信息
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivitySkuList(_activitySkuList []DrfTxdActivitySkuBo) error {
	r._activitySkuList = _activitySkuList
	r.Set("activity_sku_list", _activitySkuList)
	return nil
}

// Get ActivitySkuList Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivitySkuList() []DrfTxdActivitySkuBo {
	return r._activitySkuList
}

// Set is VersionId Setter
// 数据版本Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// Get VersionId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// Set is ActivityId Setter
// 大润发活动Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivityId() string {
	return r._activityId
}

// Set is ShopId Setter
// 活动对应的门店Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetShopId() string {
	return r._shopId
}
