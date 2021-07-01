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
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest struct {
    model.Params
    // 确认入职后入职单据请求
    _paramSyncEntryReceiptRequest   *SyncEntryReceiptRequest
}

// 初始化AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest对象
func NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest{
    return &AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.moka.entry.receipt.write"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSyncEntryReceiptRequest Setter
// 确认入职后入职单据请求
func (r *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) SetParamSyncEntryReceiptRequest(_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest) error {
    r._paramSyncEntryReceiptRequest = _paramSyncEntryReceiptRequest
    r.Set("param_sync_entry_receipt_request", _paramSyncEntryReceiptRequest)
    return nil
}

// ParamSyncEntryReceiptRequest Getter
func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest) GetParamSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
    return r._paramSyncEntryReceiptRequest
}
