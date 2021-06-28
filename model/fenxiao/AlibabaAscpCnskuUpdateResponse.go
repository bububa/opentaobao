package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台货品修改接口 APIResponse
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
type AlibabaAscpCnskuUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAscpCnskuUpdateResponse `json:"alibaba_ascp_cnsku_update_response,omitempty"` 
    AlibabaAscpCnskuUpdateResponse
}

/* model for simplify = false
type AlibabaAscpCnskuUpdateResponse struct {

    // 异常信息
    
    ErrorMessages  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_messages,omitempty"`
    

    // 货品id
    
    Data   string `json:"data,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 是否系统异常
    
    IsSystemFailed   bool `json:"is_system_failed,omitempty"`
    

    // 异常信息Code
    
    SysErrorCode   string `json:"sys_error_code,omitempty"`
    

}
*/

type AlibabaAscpCnskuUpdateResponse struct {

    // 异常信息
    ErrorMessages   []string `json:"error_messages,omitempty"`

    // 货品id
    Data   string `json:"data,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 是否系统异常
    IsSystemFailed   bool `json:"is_system_failed,omitempty"`

    // 异常信息Code
    SysErrorCode   string `json:"sys_error_code,omitempty"`

}
