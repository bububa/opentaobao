package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据菜鸟账号ID指派小件员 APIResponse
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_guoguo_cp_backup_assigncourierbyid_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 指派/改派是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"