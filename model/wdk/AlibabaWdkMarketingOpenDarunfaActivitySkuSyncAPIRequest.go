package wdk

import (
	"net/url"
	"sync"

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
	// 大润发活动Id
	_activityId string
	// 活动对应的门店Id
	_shopId string
	// 数据版本Id
	_versionId int64
}

// NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest 初始化AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest对象
func NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest() *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest {
	return &AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) Reset() {
	r._activitySkuList = r._activitySkuList[:0]
	r._activityId = ""
	r._shopId = ""
	r._versionId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.darunfa.activity.sku.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivitySkuList is ActivitySkuList Setter
// 淘鲜达活动商品信息
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivitySkuList(_activitySkuList []DrfTxdActivitySkuBo) error {
	r._activitySkuList = _activitySkuList
	r.Set("activity_sku_list", _activitySkuList)
	return nil
}

// GetActivitySkuList ActivitySkuList Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivitySkuList() []DrfTxdActivitySkuBo {
	return r._activitySkuList
}

// SetActivityId is ActivityId Setter
// 大润发活动Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetShopId is ShopId Setter
// 活动对应的门店Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetShopId() string {
	return r._shopId
}

// SetVersionId is VersionId Setter
// 数据版本Id
func (r *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) GetVersionId() int64 {
	return r._versionId
}

var poolAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest()
	},
}

// GetAlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest 从 sync.Pool 获取 AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest
func GetAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest() *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest {
	return poolAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest.Get().(*AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest)
}

// ReleaseAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest 将 AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest(v *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest.Put(v)
}
