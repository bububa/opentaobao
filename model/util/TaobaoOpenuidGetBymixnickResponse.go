package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过mixnick转换openuid APIResponse
taobao.openuid.get.bymixnick

通过mixnick转换openuid
*/
type TaobaoOpenuidGetBymixnickAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openuid_get_bymixnick_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // OpenUID
    
    OpenUid   string `json:"open_uid,omitempty" xml:"