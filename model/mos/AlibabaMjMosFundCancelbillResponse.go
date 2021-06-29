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
    AlibabaMjMosFundCancelbillResponse
}

type AlibabaMjMosFundCancelbillResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_mos_fund_cancelbill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
}
