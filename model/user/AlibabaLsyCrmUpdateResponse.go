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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_lsy_crm_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty" xml:"