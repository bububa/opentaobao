package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提货核销 APIResponse
alibaba.mj.oc.confpickupgoods

此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
*/
type AlibabaMjOcConfpickupgoodsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcConfpickupgoodsResponse `json:"alibaba_mj_oc_confpickupgoods_response,omitempty"`
}

type AlibabaMjOcConfpickupgoodsResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
