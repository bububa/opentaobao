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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_lsy_crm_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty" xml:"