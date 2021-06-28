package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
跟进客资状态接口 APIResponse
alibaba.lsy.crm.update

同步客资状态接口
*/
type AlibabaLsyCrmUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmUpdateResponse
}

type AlibabaLsyCrmUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`

    
    // 错误码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 是否成功跟进客资状态
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
}
