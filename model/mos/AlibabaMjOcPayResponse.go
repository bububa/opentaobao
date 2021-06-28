package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
POS收银成功后订单同步 APIResponse
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
*/
type AlibabaMjOcPayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_pay_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errCode
    
    ExCode   int64 `json:"ex_code,omitempty" xml:"