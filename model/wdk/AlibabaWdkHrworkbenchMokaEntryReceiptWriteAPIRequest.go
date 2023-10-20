package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest 摩卡确认入职后往入职单据表写数据接口 API请求
// alibaba.wdk.hrworkbench.moka.entry.receipt.write
//
// 摩卡确认入职后往入职单据表写数据接口
type AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest struct {
	model.Params
	// 确认入职后入职单据请求
	_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest
}

// NewAlibabawdkhrworkbenchmokaentryreceiptwriteRequest 初始化AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest对象
func NewAlibabawdkhrworkbenchmokaentryreceiptwriteRequest() *AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest {
	return &AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.hrworkbench.moka.entry.receipt.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncEntryReceiptRequest is ParamSyncEntryReceiptRequest Setter
// 确认入职后入职单据请求
func (r *AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest) SetParamSyncEntryReceiptRequest(_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest) error {
	r._paramSyncEntryReceiptRequest = _paramSyncEntryReceiptRequest
	r.Set("param_sync_entry_receipt_request", _paramSyncEntryReceiptRequest)
	return nil
}

// GetParamSyncEntryReceiptRequest ParamSyncEntryReceiptRequest Getter
func (r AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest) GetParamSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
	return r._paramSyncEntryReceiptRequest
}
