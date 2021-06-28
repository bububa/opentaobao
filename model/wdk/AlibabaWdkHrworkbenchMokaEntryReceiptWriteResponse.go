package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
摩卡确认入职后往入职单据表写数据接口 APIResponse
alibaba.wdk.hrworkbench.moka.entry.receipt.write

摩卡确认入职后往入职单据表写数据接口
*/
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_hrworkbench_moka_entry_receipt_write_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   bool `json:"data,omitempty" xml:"