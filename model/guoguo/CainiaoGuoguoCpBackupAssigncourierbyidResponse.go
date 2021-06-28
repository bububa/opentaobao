package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据菜鸟账号ID指派小件员 APIResponse
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoGuoguoCpBackupAssigncourierbyidResponse `json:"cainiao_guoguo_cp_backup_assigncourierbyid_response,omitempty"` 
    CainiaoGuoguoCpBackupAssigncourierbyidResponse
}

/* model for simplify = false
type CainiaoGuoguoCpBackupAssigncourierbyidResponse struct {

    // 指派/改派是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 错误信息描述
    
    StatusMessage   string `json:"status_message,omitempty"`
    

}
*/

type CainiaoGuoguoCpBackupAssigncourierbyidResponse struct {

    // 指派/改派是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    StatusCode   string `json:"status_code,omitempty"`

    // 错误信息描述
    StatusMessage   string `json:"status_message,omitempty"`

}
