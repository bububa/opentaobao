package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摩卡确认入职后往入职单据表写数据接口 API请求
alibaba.wdk.hrworkbench.moka.entry.receipt.write

摩卡确认入职后往入职单据表写数据接口
*/
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest struct {
    model.Params
    // 确认入职后入职单据请求
    _paramSyncEntryReceiptRequest   *SyncEntryReceiptRequest
}

// 初始化AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest对象
func NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest{
    return &AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.moka.entry.receipt.write"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSyncEntryReceiptRequest Setter
// 确认入职后入职单据请求
func (r *AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) SetParamSyncEntryReceiptRequest(_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest) error {
    r._paramSyncEntryReceiptRequest = _paramSyncEntryReceiptRequest
    r.Set("param_sync_entry_receipt_request", _paramSyncEntryReceiptRequest)
    return nil
}

// ParamSyncEntryReceiptRequest Getter
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetParamSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
    return r._paramSyncEntryReceiptRequest
}
