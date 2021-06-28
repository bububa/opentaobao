package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消付款单 APIResponse
alibaba.mj.mos.fund.cancelbill

取消付款单
*/
type AlibabaMjMosFundCancelbillAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_mos_fund_cancelbill_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   bool `json:"data,omitempty" xml:"