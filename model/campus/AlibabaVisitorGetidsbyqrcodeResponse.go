package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据访客二维码查访客行程id APIResponse
alibaba.visitor.getidsbyqrcode

根据支付宝阿里访客小程序的动态二维码查询来访行程id
*/
type AlibabaVisitorGetidsbyqrcodeAPIResponse struct {
    model.CommonResponse
    AlibabaVisitorGetidsbyqrcodeResponse
}

type AlibabaVisitorGetidsbyqrcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_visitor_getidsbyqrcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
