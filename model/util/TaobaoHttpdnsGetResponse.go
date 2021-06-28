package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
TOPDNS配置 APIResponse
taobao.httpdns.get

获取TOP DNS配置
*/
type TaobaoHttpdnsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"httpdns_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // HTTP DNS配置信息
    
    Result   string `json:"result,omitempty" xml:"