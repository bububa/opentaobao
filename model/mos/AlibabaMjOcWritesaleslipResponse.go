package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开票占库 APIResponse
alibaba.mj.oc.writesaleslip

开票占库
*/
type AlibabaMjOcWritesaleslipAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_writesaleslip_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"