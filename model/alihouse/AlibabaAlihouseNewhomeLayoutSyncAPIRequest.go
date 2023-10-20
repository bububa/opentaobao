package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomelayoutsyncAPIRequest 房通户型数据同步 API请求
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
type AlibabaalihousenewhomelayoutsyncAPIRequest struct {
	model.Params
	// 数据
	_syncProjectLayoutData *SyncProjectLayoutDto
}

// NewAlibabaalihousenewhomelayoutsyncRequest 初始化AlibabaalihousenewhomelayoutsyncAPIRequest对象
func NewAlibabaalihousenewhomelayoutsyncRequest() *AlibabaalihousenewhomelayoutsyncAPIRequest {
	return &AlibabaalihousenewhomelayoutsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomelayoutsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.layout.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomelayoutsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomelayoutsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProjectLayoutData is SyncProjectLayoutData Setter
// 数据
func (r *AlibabaalihousenewhomelayoutsyncAPIRequest) SetSyncProjectLayoutData(_syncProjectLayoutData *SyncProjectLayoutDto) error {
	r._syncProjectLayoutData = _syncProjectLayoutData
	r.Set("sync_project_layout_data", _syncProjectLayoutData)
	return nil
}

// GetSyncProjectLayoutData SyncProjectLayoutData Getter
func (r AlibabaalihousenewhomelayoutsyncAPIRequest) GetSyncProjectLayoutData() *SyncProjectLayoutDto {
	return r._syncProjectLayoutData
}
