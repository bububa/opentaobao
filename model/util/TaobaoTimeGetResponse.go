package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取淘宝系统当前时间 APIResponse
taobao.time.get

获取淘宝系统当前时间
*/
type TaobaoTimeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"time_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
    
    Time   string `json:"time,omitempty" xml:"