package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建一个付款单 APIResponse
alibaba.mj.mos.fund.createbill

创建一个付款单
*/
type AlibabaMjMosFundCreatebillAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMosFundCreatebillResponse `json:"alibaba_mj_mos_fund_createbill_response,omitempty"`
}

type AlibabaMjMosFundCreatebillResponse struct {

    // data
    Data   string `json:"data,omitempty"`

}
