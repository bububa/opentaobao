package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建客资 APIResponse
alibaba.lsy.crm.create

欢客调用该接口进行客资创建
*/
type AlibabaLsyCrmCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaLsyCrmCreateResponse `json:"alibaba_lsy_crm_create_response,omitempty"`
}

type AlibabaLsyCrmCreateResponse struct {

    // 错误提示
    FailMsg   string `json:"fail_msg,omitempty"`

    // 错误码
    FailCode   string `json:"fail_code,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

    // 返回的数据
    Data   *NrtCreateRecordReturnDTO `json:"data,omitempty"`

}
