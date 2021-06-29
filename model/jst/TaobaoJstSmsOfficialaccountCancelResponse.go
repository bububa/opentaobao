package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔取消公众号订购 APIResponse
taobao.jst.sms.officialaccount.cancel

聚石塔取消公众号订购
*/
type TaobaoJstSmsOfficialaccountCancelAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsOfficialaccountCancelResponse
}

type TaobaoJstSmsOfficialaccountCancelResponse struct {
    XMLName xml.Name `xml:"jst_sms_officialaccount_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统异常
    
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`

    
    // 成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
    // 请求id
    
    ResponseId   string `json:"response_id,omitempty" xml:"response_id,omitempty"`

    
    // 成功
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
    // 系统异常
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
