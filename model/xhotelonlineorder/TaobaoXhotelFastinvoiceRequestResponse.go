package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求回传 APIResponse
taobao.xhotel.fastinvoice.request

极速开票开票请求回传,用于记录航信开票请求数据
*/
type TaobaoXhotelFastinvoiceRequestAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelFastinvoiceRequestResponse
}

type TaobaoXhotelFastinvoiceRequestResponse struct {
    XMLName xml.Name `xml:"xhotel_fastinvoice_request_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`

    
    // errorCode
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // errorMsg
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
