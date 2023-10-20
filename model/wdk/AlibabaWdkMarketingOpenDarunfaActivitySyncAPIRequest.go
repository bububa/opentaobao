package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopendarunfaactivitysyncAPIRequest 活动数据同步 API请求
// alibaba.wdk.marketing.open.darunfa.activity.sync
//
// 大润发活动数据同步
type AlibabawdkmarketingopendarunfaactivitysyncAPIRequest struct {
	model.Params
	// 大润发活动数据
	_activityList []DrfTxdActivityBo
	// 门店Id
	_shopId string
	// 版本ID
	_versionId int64
}

// NewAlibabawdkmarketingopendarunfaactivitysyncRequest 初始化AlibabawdkmarketingopendarunfaactivitysyncAPIRequest对象
func NewAlibabawdkmarketingopendarunfaactivitysyncRequest() *AlibabawdkmarketingopendarunfaactivitysyncAPIRequest {
	return &AlibabawdkmarketingopendarunfaactivitysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.darunfa.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityList is ActivityList Setter
// 大润发活动数据
func (r *AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) SetActivityList(_activityList []DrfTxdActivityBo) error {
	r._activityList = _activityList
	r.Set("activity_list", _activityList)
	return nil
}

// GetActivityList ActivityList Getter
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetActivityList() []DrfTxdActivityBo {
	return r._activityList
}

// SetShopId is ShopId Setter
// 门店Id
func (r *AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetShopId() string {
	return r._shopId
}

// SetVersionId is VersionId Setter
// 版本ID
func (r *AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r AlibabawdkmarketingopendarunfaactivitysyncAPIRequest) GetVersionId() int64 {
	return r._versionId
}
