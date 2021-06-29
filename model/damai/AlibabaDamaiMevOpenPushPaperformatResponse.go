package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat APIResponse
alibaba.damai.mev.open.push.paperformat

pushPaperFormat
*/
type AlibabaDamaiMevOpenPushPaperformatAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushPaperformatResponse
}

type AlibabaDamaiMevOpenPushPaperformatResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_push_paperformat_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushPaperformatResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
