package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个付款单 APIResponse
alibaba.mj.mos.fund.createbill

创建一个付款单
*/
type AlibabaMjMosFundCreatebillAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_mos_fund_createbill_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   string `json:"data,omitempty" xml:"