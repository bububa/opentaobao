package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提货核销 APIResponse
alibaba.mj.oc.confpickupgoods

此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
*/
type AlibabaMjOcConfpickupgoodsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_confpickupgoods_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"