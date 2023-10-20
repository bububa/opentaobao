package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenVersionApplyAPIRequest 数据同步版本号申请 API请求
// alibaba.wdk.marketing.open.version.apply
//
// 数据同步版本号申请
type AlibabaWdkMarketingOpenVersionApplyAPIRequest struct {
	model.Params
	// 同步版本信息
	_syncVersion *SyncVersionBo
}

// NewAlibabaWdkMarketingOpenVersionApplyRequest 初始化AlibabaWdkMarketingOpenVersionApplyAPIRequest对象
func NewAlibabaWdkMarketingOpenVersionApplyRequest() *AlibabaWdkMarketingOpenVersionApplyAPIRequest {
	return &AlibabaWdkMarketingOpenVersionApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingOpenVersionApplyAPIRequest) Reset() {
	r._syncVersion = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.version.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncVersion is SyncVersion Setter
// 同步版本信息
func (r *AlibabaWdkMarketingOpenVersionApplyAPIRequest) SetSyncVersion(_syncVersion *SyncVersionBo) error {
	r._syncVersion = _syncVersion
	r.Set("sync_version", _syncVersion)
	return nil
}

// GetSyncVersion SyncVersion Getter
func (r AlibabaWdkMarketingOpenVersionApplyAPIRequest) GetSyncVersion() *SyncVersionBo {
	return r._syncVersion
}

var poolAlibabaWdkMarketingOpenVersionApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingOpenVersionApplyRequest()
	},
}

// GetAlibabaWdkMarketingOpenVersionApplyRequest 从 sync.Pool 获取 AlibabaWdkMarketingOpenVersionApplyAPIRequest
func GetAlibabaWdkMarketingOpenVersionApplyAPIRequest() *AlibabaWdkMarketingOpenVersionApplyAPIRequest {
	return poolAlibabaWdkMarketingOpenVersionApplyAPIRequest.Get().(*AlibabaWdkMarketingOpenVersionApplyAPIRequest)
}

// ReleaseAlibabaWdkMarketingOpenVersionApplyAPIRequest 将 AlibabaWdkMarketingOpenVersionApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingOpenVersionApplyAPIRequest(v *AlibabaWdkMarketingOpenVersionApplyAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingOpenVersionApplyAPIRequest.Put(v)
}
