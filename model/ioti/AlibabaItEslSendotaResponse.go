package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签ota接口 APIResponse
alibaba.it.esl.sendota

厂测接口，电子价签ota接口
*/
type AlibabaItEslSendotaAPIResponse struct {
    model.CommonResponse
    AlibabaItEslSendotaResponse
}

type AlibabaItEslSendotaResponse struct {
    XMLName xml.Name `xml:"alibaba_it_esl_sendota_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Can not find Corresponding AP MAC with ESL
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
