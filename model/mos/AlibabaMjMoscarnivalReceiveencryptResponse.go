package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据加密手机号领券 APIResponse
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券
*/
type AlibabaMjMoscarnivalReceiveencryptAPIResponse struct {
    model.CommonResponse
    AlibabaMjMoscarnivalReceiveencryptResponse
}

type AlibabaMjMoscarnivalReceiveencryptResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_moscarnival_receiveencrypt_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaMjMoscarnivalReceiveencryptResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
