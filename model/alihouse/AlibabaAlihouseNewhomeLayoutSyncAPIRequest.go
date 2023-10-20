package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLayoutSyncAPIRequest 房通户型数据同步 API请求
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
type AlibabaAlihouseNewhomeLayoutSyncAPIRequest struct {
	model.Params
	// 数据
	_syncProjectLayoutData *SyncProjectLayoutDto
}

// NewAlibabaAlihouseNewhomeLayoutSyncRequest 初始化AlibabaAlihouseNewhomeLayoutSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeLayoutSyncRequest() *AlibabaAlihouseNewhomeLayoutSyncAPIRequest {
	return &AlibabaAlihouseNewhomeLayoutSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeLayoutSyncAPIRequest) Reset() {
	r._syncProjectLayoutData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.layout.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProjectLayoutData is SyncProjectLayoutData Setter
// 数据
func (r *AlibabaAlihouseNewhomeLayoutSyncAPIRequest) SetSyncProjectLayoutData(_syncProjectLayoutData *SyncProjectLayoutDto) error {
	r._syncProjectLayoutData = _syncProjectLayoutData
	r.Set("sync_project_layout_data", _syncProjectLayoutData)
	return nil
}

// GetSyncProjectLayoutData SyncProjectLayoutData Getter
func (r AlibabaAlihouseNewhomeLayoutSyncAPIRequest) GetSyncProjectLayoutData() *SyncProjectLayoutDto {
	return r._syncProjectLayoutData
}

var poolAlibabaAlihouseNewhomeLayoutSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeLayoutSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeLayoutSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeLayoutSyncAPIRequest
func GetAlibabaAlihouseNewhomeLayoutSyncAPIRequest() *AlibabaAlihouseNewhomeLayoutSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeLayoutSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeLayoutSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeLayoutSyncAPIRequest 将 AlibabaAlihouseNewhomeLayoutSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLayoutSyncAPIRequest(v *AlibabaAlihouseNewhomeLayoutSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLayoutSyncAPIRequest.Put(v)
}
