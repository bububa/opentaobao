package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口batchPushTicket APIResponse
alibaba.damai.mev.open.batchpushticket

批量推送票单
*/
type AlibabaDamaiMevOpenBatchpushticketAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenBatchpushticketResponse
}

type AlibabaDamaiMevOpenBatchpushticketResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_batchpushticket_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenBatchpushticketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
