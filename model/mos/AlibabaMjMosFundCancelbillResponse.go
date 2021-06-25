package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消付款单 APIResponse
alibaba.mj.mos.fund.cancelbill

取消付款单
*/
type AlibabaMjMosFundCancelbillAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMosFundCancelbillResponse `json:"alibaba_mj_mos_fund_cancelbill_response,omitempty"`
}

type AlibabaMjMosFundCancelbillResponse struct {

    // data
    Data   bool `json:"data,omitempty"`

}
