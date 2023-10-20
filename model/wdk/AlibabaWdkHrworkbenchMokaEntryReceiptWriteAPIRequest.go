package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest 摩卡确认入职后往入职单据表写数据接口 API请求
// alibaba.wdk.hrworkbench.moka.entry.receipt.write
//
// 摩卡确认入职后往入职单据表写数据接口
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest struct {
	model.Params
	// 确认入职后入职单据请求
	_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest
}

// NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest 初始化AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest对象
func NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest {
	return &AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) Reset() {
	r._paramSyncEntryReceiptRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.hrworkbench.moka.entry.receipt.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncEntryReceiptRequest is ParamSyncEntryReceiptRequest Setter
// 确认入职后入职单据请求
func (r *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) SetParamSyncEntryReceiptRequest(_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest) error {
	r._paramSyncEntryReceiptRequest = _paramSyncEntryReceiptRequest
	r.Set("param_sync_entry_receipt_request", _paramSyncEntryReceiptRequest)
	return nil
}

// GetParamSyncEntryReceiptRequest ParamSyncEntryReceiptRequest Getter
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetParamSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
	return r._paramSyncEntryReceiptRequest
}

var poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest()
	},
}

// GetAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest 从 sync.Pool 获取 AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest
func GetAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest {
	return poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest.Get().(*AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest)
}

// ReleaseAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest 将 AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest(v *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) {
	v.Reset()
	poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest.Put(v)
}
