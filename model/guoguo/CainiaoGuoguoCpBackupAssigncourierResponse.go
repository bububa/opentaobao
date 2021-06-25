package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CP兜底后指定接单的小件员 APIResponse
cainiao.guoguo.cp.backup.assigncourier

CP兜底后指定接单的小件员；CP改派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierAPIResponse struct {
    model.CommonResponse
    Response *CainiaoGuoguoCpBackupAssigncourierResponse `json:"cainiao_guoguo_cp_backup_assigncourier_response,omitempty"`
}

type CainiaoGuoguoCpBackupAssigncourierResponse struct {

    // 错误码
    StatusCode   string `json:"status_code,omitempty"`

    // 错误信息
    StatusMessage   string `json:"status_message,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
