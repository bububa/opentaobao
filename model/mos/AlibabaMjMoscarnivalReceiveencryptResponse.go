package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据加密手机号领券 API返回值 
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券
*/
type AlibabaMjMoscarnivalReceiveencryptAPIResponse struct {
    model.CommonResponse
    AlibabaMjMoscarnivalReceiveencryptResponse
}

// 根据加密手机号领券 成功返回结果
type AlibabaMjMoscarnivalReceiveencryptResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_moscarnival_receiveencrypt_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaMjMoscarnivalReceiveencryptResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
