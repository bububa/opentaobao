package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建客资 APIResponse
alibaba.lsy.crm.create

欢客调用该接口进行客资创建
*/
type AlibabaLsyCrmCreateAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmCreateResponse
}

type AlibabaLsyCrmCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`

    
    // 错误码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 返回的数据
    
    Data   *NrtCreateRecordReturnDTO `json:"data,omitempty" xml:"data,omitempty"`

    
}
