package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断支付宝用户是否绑定淘宝账号 APIResponse
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowAlipayIsbindingtbaccountResponse `json:"alibaba_aliqin_flow_alipay_isbindingtbaccount_response,omitempty"`
}

type AlibabaAliqinFlowAlipayIsbindingtbaccountResponse struct {

    // error
    Error   bool `json:"error,omitempty"`

    // TRUE代表绑定，FALSE代表未绑定
    Value   string `json:"value,omitempty"`

    // errorCode
    AlicomFlowErrorCode   string `json:"alicom_flow_error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
