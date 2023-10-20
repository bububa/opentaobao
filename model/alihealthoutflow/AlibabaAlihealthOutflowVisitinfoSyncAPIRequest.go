package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowvisitinfosyncAPIRequest 处方外流-问诊、处方同步 API请求
// alibaba.alihealth.outflow.visitinfo.sync
//
// 阿里健康-处方外流-对外提供问诊、处方功能
type AlibabaalihealthoutflowvisitinfosyncAPIRequest struct {
	model.Params
	// 入参
	_syncVisitInfoRequest *SyncVisitInfoRequest
}

// NewAlibabaalihealthoutflowvisitinfosyncRequest 初始化AlibabaalihealthoutflowvisitinfosyncAPIRequest对象
func NewAlibabaalihealthoutflowvisitinfosyncRequest() *AlibabaalihealthoutflowvisitinfosyncAPIRequest {
	return &AlibabaalihealthoutflowvisitinfosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowvisitinfosyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.visitinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowvisitinfosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowvisitinfosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncVisitInfoRequest is SyncVisitInfoRequest Setter
// 入参
func (r *AlibabaalihealthoutflowvisitinfosyncAPIRequest) SetSyncVisitInfoRequest(_syncVisitInfoRequest *SyncVisitInfoRequest) error {
	r._syncVisitInfoRequest = _syncVisitInfoRequest
	r.Set("sync_visit_info_request", _syncVisitInfoRequest)
	return nil
}

// GetSyncVisitInfoRequest SyncVisitInfoRequest Getter
func (r AlibabaalihealthoutflowvisitinfosyncAPIRequest) GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
	return r._syncVisitInfoRequest
}
