package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开票占库 APIResponse
alibaba.mj.oc.writesaleslip

开票占库
*/
type AlibabaMjOcWritesaleslipAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcWritesaleslipResponse `json:"alibaba_mj_oc_writesaleslip_response,omitempty"`
}

type AlibabaMjOcWritesaleslipResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
