package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowOperationinfoSyncAPIRequest 处方外流-操作信息同步 API请求
// alibaba.alihealth.outflow.operationinfo.sync
//
// 阿里健康-处方外流-对外提供同步操作信息功能
type AlibabaAlihealthOutflowOperationinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncOperationInfoRequest *SyncOperationInfoRequest
}

// NewAlibabaAlihealthOutflowOperationinfoSyncRequest 初始化AlibabaAlihealthOutflowOperationinfoSyncAPIRequest对象
func NewAlibabaAlihealthOutflowOperationinfoSyncRequest() *AlibabaAlihealthOutflowOperationinfoSyncAPIRequest {
	return &AlibabaAlihealthOutflowOperationinfoSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) Reset() {
	r._syncOperationInfoRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.operationinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncOperationInfoRequest is SyncOperationInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) SetSyncOperationInfoRequest(_syncOperationInfoRequest *SyncOperationInfoRequest) error {
	r._syncOperationInfoRequest = _syncOperationInfoRequest
	r.Set("sync_operation_info_request", _syncOperationInfoRequest)
	return nil
}

// GetSyncOperationInfoRequest SyncOperationInfoRequest Getter
func (r AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) GetSyncOperationInfoRequest() *SyncOperationInfoRequest {
	return r._syncOperationInfoRequest
}

var poolAlibabaAlihealthOutflowOperationinfoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowOperationinfoSyncRequest()
	},
}

// GetAlibabaAlihealthOutflowOperationinfoSyncRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowOperationinfoSyncAPIRequest
func GetAlibabaAlihealthOutflowOperationinfoSyncAPIRequest() *AlibabaAlihealthOutflowOperationinfoSyncAPIRequest {
	return poolAlibabaAlihealthOutflowOperationinfoSyncAPIRequest.Get().(*AlibabaAlihealthOutflowOperationinfoSyncAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowOperationinfoSyncAPIRequest 将 AlibabaAlihealthOutflowOperationinfoSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowOperationinfoSyncAPIRequest(v *AlibabaAlihealthOutflowOperationinfoSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowOperationinfoSyncAPIRequest.Put(v)
}
