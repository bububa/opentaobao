package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口pushStand API返回值 
alibaba.damai.mev.open.pushstand

pushStand
*/
type AlibabaDamaiMevOpenPushstandAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushstandAPIResponseModel
}

// 大麦换验平台-第三方对外开放-看台接口pushStand 成功返回结果
type AlibabaDamaiMevOpenPushstandAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushstand_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaDamaiMevOpenPushstandResult `json:"result,omitempty" xml:"result,omitempty"`
}
