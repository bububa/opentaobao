package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口invalidTicket API返回值 
alibaba.damai.mev.open.invalidticket

开放接口 使票无效
*/
type AlibabaDamaiMevOpenInvalidticketAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenInvalidticketAPIResponseModel
}

// 大麦换验平台-第三方对外开放-票单接口invalidTicket 成功返回结果
type AlibabaDamaiMevOpenInvalidticketAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_invalidticket_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaDamaiMevOpenInvalidticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
