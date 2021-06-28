package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群设置管理员 APIResponse
taobao.openim.tribe.setmanager

OPENIM群设置管理员
*/
type TaobaoOpenimTribeSetmanagerAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_tribe_setmanager_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"