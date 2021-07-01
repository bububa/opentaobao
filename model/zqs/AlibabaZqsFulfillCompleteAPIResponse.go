package zqs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
周期购履约完成接口 API返回值 
alibaba.zqs.fulfill.complete

周期购履约完成接口
*/
type AlibabaZqsFulfillCompleteAPIResponse struct {
    model.CommonResponse
    AlibabaZqsFulfillCompleteAPIResponseModel
}

// 周期购履约完成接口 成功返回结果
type AlibabaZqsFulfillCompleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_zqs_fulfill_complete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异常描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 异常code
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 执行结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
