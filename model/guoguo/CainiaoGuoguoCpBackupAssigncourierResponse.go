package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CP兜底后指定接单的小件员 APIResponse
cainiao.guoguo.cp.backup.assigncourier

CP兜底后指定接单的小件员；CP改派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_guoguo_cp_backup_assigncourier_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    StatusCode   string `json:"status_code,omitempty" xml:"