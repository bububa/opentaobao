package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款详情 API返回值 
alibaba.wdk.reverse.reversedetail

退款详情
*/
type AlibabaWdkReverseReversedetailAPIResponse struct {
    model.CommonResponse
    AlibabaWdkReverseReversedetailAPIResponseModel
}

// 退款详情 成功返回结果
type AlibabaWdkReverseReversedetailAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_reverse_reversedetail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
