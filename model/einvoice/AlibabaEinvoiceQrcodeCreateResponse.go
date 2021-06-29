package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票二维码生成 APIResponse
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
*/
type AlibabaEinvoiceQrcodeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceQrcodeCreateResponse
}

type AlibabaEinvoiceQrcodeCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_qrcode_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaEinvoiceQrcodeCreateResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
