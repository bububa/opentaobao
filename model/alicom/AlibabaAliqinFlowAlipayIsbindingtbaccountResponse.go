package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断支付宝用户是否绑定淘宝账号 APIResponse
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_alipay_isbindingtbaccount_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // error
    
    Error   bool `json:"error,omitempty" xml:"