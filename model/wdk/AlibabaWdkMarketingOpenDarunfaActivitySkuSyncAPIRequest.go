package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest 营销商品数据同步 API请求
// alibaba.wdk.marketing.open.darunfa.activity.sku.sync
//
// 大润发营销商品数据同步
type AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest struct {
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

// NewAlibabawdkmarketingopendarunfaactivityskusyncRequest 初始化AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest对象
func NewAlibabawdkmarketingopendarunfaactivityskusyncRequest() *AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest {
	return &AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.darunfa.activity.sku.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivitySkuList is ActivitySkuList Setter
// 淘鲜达活动商品信息
func (r *AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) SetActivitySkuList(_activitySkuList []DrfTxdActivitySkuBo) error {
	r._activitySkuList = _activitySkuList
	r.Set("activity_sku_list", _activitySkuList)
	return nil
}

// GetActivitySkuList ActivitySkuList Getter
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetActivitySkuList() []DrfTxdActivitySkuBo {
	return r._activitySkuList
}

// SetActivityId is ActivityId Setter
// 大润发活动Id
func (r *AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetShopId is ShopId Setter
// 活动对应的门店Id
func (r *AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetShopId() string {
	return r._shopId
}

// SetVersionId is VersionId Setter
// 数据版本Id
func (r *AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest) GetVersionId() int64 {
	return r._versionId
}
