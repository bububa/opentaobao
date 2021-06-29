package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求完成 APIResponse
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据
*/
type TaobaoXhotelFastinvoiceCompleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelFastinvoiceCompleteResponse
}

type TaobaoXhotelFastinvoiceCompleteResponse struct {
    XMLName xml.Name `xml:"xhotel_fastinvoice_complete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`

    
    // errorCode
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // errorMsg
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
