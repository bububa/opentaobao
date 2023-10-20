package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopenversionapplyAPIRequest 数据同步版本号申请 API请求
// alibaba.wdk.marketing.open.version.apply
//
// 数据同步版本号申请
type AlibabawdkmarketingopenversionapplyAPIRequest struct {
	model.Params
	// 同步版本信息
	_syncVersion *SyncVersionBo
}

// NewAlibabawdkmarketingopenversionapplyRequest 初始化AlibabawdkmarketingopenversionapplyAPIRequest对象
func NewAlibabawdkmarketingopenversionapplyRequest() *AlibabawdkmarketingopenversionapplyAPIRequest {
	return &AlibabawdkmarketingopenversionapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopenversionapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.version.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopenversionapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopenversionapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncVersion is SyncVersion Setter
// 同步版本信息
func (r *AlibabawdkmarketingopenversionapplyAPIRequest) SetSyncVersion(_syncVersion *SyncVersionBo) error {
	r._syncVersion = _syncVersion
	r.Set("sync_version", _syncVersion)
	return nil
}

// GetSyncVersion SyncVersion Getter
func (r AlibabawdkmarketingopenversionapplyAPIRequest) GetSyncVersion() *SyncVersionBo {
	return r._syncVersion
}
