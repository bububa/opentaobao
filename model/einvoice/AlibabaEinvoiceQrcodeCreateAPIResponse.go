package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票二维码生成 API返回值 
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
*/
type AlibabaEinvoiceQrcodeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceQrcodeCreateAPIResponseModel
}

// 扫码开票二维码生成 成功返回结果
type AlibabaEinvoiceQrcodeCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_qrcode_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaEinvoiceQrcodeCreateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
