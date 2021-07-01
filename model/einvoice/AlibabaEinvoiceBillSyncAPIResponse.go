package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单同步 API返回值 
alibaba.einvoice.bill.sync

电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票
*/
type AlibabaEinvoiceBillSyncAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceBillSyncAPIResponseModel
}

// 结算单同步 成功返回结果
type AlibabaEinvoiceBillSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_bill_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 插入操作返回二维码
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 返回码
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // success
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
