package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenVersionApplyAPIRequest
数据同步版本号申请 API请求
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请 */
type AlibabaWdkMarketingOpenVersionApplyAPIRequest struct {
	model.Params
	// 同步版本信息
	_syncVersion *SyncVersionBo
}

// NewAlibabaWdkMarketingOpenVersionApplyRequest 初始化AlibabaWdkMarketingOpenVersionApplyAPIRequest对象
func NewAlibabaWdkMarketingOpenVersionApplyRequest() *AlibabaWdkMarketingOpenVersionApplyAPIRequest {
	return &AlibabaWdkMarketingOpenVersionApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.version.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SyncVersion Setter
// 同步版本信息
func (r *AlibabaWdkMarketingOpenVersionApplyAPIRequest) SetSyncVersion(_syncVersion *SyncVersionBo) error {
	r._syncVersion = _syncVersion
	r.Set("sync_version", _syncVersion)
	return nil
}

// Get SyncVersion Getter
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetSyncVersion() *SyncVersionBo {
	return r._syncVersion
}
