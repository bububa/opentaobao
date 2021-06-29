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
    TaobaoHttpdnsGetResponse
}

type TaobaoHttpdnsGetResponse struct {
    XMLName xml.Name `xml:"httpdns_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // HTTP DNS配置信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
