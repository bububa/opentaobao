package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowVisitinfoSyncAPIRequest 处方外流-问诊、处方同步 API请求
// alibaba.alihealth.outflow.visitinfo.sync
//
// 阿里健康-处方外流-对外提供问诊、处方功能
type AlibabaAlihealthOutflowVisitinfoSyncAPIRequest struct {
	model.Params
	// 入参
	_syncVisitInfoRequest *SyncVisitInfoRequest
}

// NewAlibabaAlihealthOutflowVisitinfoSyncRequest 初始化AlibabaAlihealthOutflowVisitinfoSyncAPIRequest对象
func NewAlibabaAlihealthOutflowVisitinfoSyncRequest() *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest {
	return &AlibabaAlihealthOutflowVisitinfoSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) Reset() {
	r._syncVisitInfoRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.visitinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncVisitInfoRequest is SyncVisitInfoRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) SetSyncVisitInfoRequest(_syncVisitInfoRequest *SyncVisitInfoRequest) error {
	r._syncVisitInfoRequest = _syncVisitInfoRequest
	r.Set("sync_visit_info_request", _syncVisitInfoRequest)
	return nil
}

// GetSyncVisitInfoRequest SyncVisitInfoRequest Getter
func (r AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
	return r._syncVisitInfoRequest
}

var poolAlibabaAlihealthOutflowVisitinfoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowVisitinfoSyncRequest()
	},
}

// GetAlibabaAlihealthOutflowVisitinfoSyncRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowVisitinfoSyncAPIRequest
func GetAlibabaAlihealthOutflowVisitinfoSyncAPIRequest() *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest {
	return poolAlibabaAlihealthOutflowVisitinfoSyncAPIRequest.Get().(*AlibabaAlihealthOutflowVisitinfoSyncAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowVisitinfoSyncAPIRequest 将 AlibabaAlihealthOutflowVisitinfoSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowVisitinfoSyncAPIRequest(v *AlibabaAlihealthOutflowVisitinfoSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowVisitinfoSyncAPIRequest.Put(v)
}
