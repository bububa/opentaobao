package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求完成 API返回值 
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据
*/
type TaobaoXhotelFastinvoiceCompleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelFastinvoiceCompleteResponse
}

// 极速开票开票请求完成 成功返回结果
type TaobaoXhotelFastinvoiceCompleteResponse struct {
    XMLName xml.Name `xml:"xhotel_fastinvoice_complete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
    // errorCode
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`
    // errorMsg
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
