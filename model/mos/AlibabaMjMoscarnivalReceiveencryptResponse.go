package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据加密手机号领券 APIResponse
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券
*/
type AlibabaMjMoscarnivalReceiveencryptAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMoscarnivalReceiveencryptResponse `json:"alibaba_mj_moscarnival_receiveencrypt_response,omitempty"`
}

type AlibabaMjMoscarnivalReceiveencryptResponse struct {

    // 返回结果
    Result   *AlibabaMjMoscarnivalReceiveencryptResultDo `json:"result,omitempty"`

}
