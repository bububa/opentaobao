package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摩卡确认入职后往入职单据表写数据接口 APIRequest
alibaba.wdk.hrworkbench.moka.entry.receipt.write

摩卡确认入职后往入职单据表写数据接口
*/
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest struct {
    model.Params

    // 确认入职后入职单据请求
    paramSyncEntryReceiptRequest   *SyncEntryReceiptRequest 

}

func NewAlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest{
    return &AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.moka.entry.receipt.write"
}

func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) SetParamSyncEntryReceiptRequest(paramSyncEntryReceiptRequest *SyncEntryReceiptRequest) error {
    r.paramSyncEntryReceiptRequest = paramSyncEntryReceiptRequest
    r.Set("param_sync_entry_receipt_request", paramSyncEntryReceiptRequest)
    return nil
}

func (r AlibabaWdkHrworkbenchMokaEntryReceiptWriteRequest) GetParamSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
    return r.paramSyncEntryReceiptRequest
}

