package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest 处方外流-处方状态同步 API请求
// alibaba.alihealth.outflow.prescription.syncstatus
//
// 阿里健康-处方外流-对外提供同步处方状态功能
type AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest struct {
	model.Params
	// 入参
	_syncStatusRequest *SyncPrescriptionStatusRequest
}

// NewAlibabaalihealthoutflowprescriptionsyncstatusRequest 初始化AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest对象
func NewAlibabaalihealthoutflowprescriptionsyncstatusRequest() *AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest {
	return &AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.syncstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncStatusRequest is SyncStatusRequest Setter
// 入参
func (r *AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest) SetSyncStatusRequest(_syncStatusRequest *SyncPrescriptionStatusRequest) error {
	r._syncStatusRequest = _syncStatusRequest
	r.Set("sync_status_request", _syncStatusRequest)
	return nil
}

// GetSyncStatusRequest SyncStatusRequest Getter
func (r AlibabaalihealthoutflowprescriptionsyncstatusAPIRequest) GetSyncStatusRequest() *SyncPrescriptionStatusRequest {
	return r._syncStatusRequest
}
