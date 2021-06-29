package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票删除 APIResponse
alitrip.btrip.invoice.setting.delete

发票删除
*/
type AlitripBtripInvoiceSettingDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripInvoiceSettingDeleteResponse
}

type AlitripBtripInvoiceSettingDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
