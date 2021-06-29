package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据菜鸟账号ID指派小件员 API返回值 
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse struct {
    model.CommonResponse
    CainiaoGuoguoCpBackupAssigncourierbyidResponse
}

// 根据菜鸟账号ID指派小件员 成功返回结果
type CainiaoGuoguoCpBackupAssigncourierbyidResponse struct {
    XMLName xml.Name `xml:"cainiao_guoguo_cp_backup_assigncourierbyid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 指派/改派是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误码
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`
    // 错误信息描述
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`
}
