package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
厂测LED控制 APIResponse
alibaba.it.esl.sendled

针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
*/
type AlibabaItEslSendledAPIResponse struct {
    model.CommonResponse
    AlibabaItEslSendledResponse
}

type AlibabaItEslSendledResponse struct {
    XMLName xml.Name `xml:"alibaba_it_esl_sendled_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Can not find Corresponding AP MAC with ESL
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
