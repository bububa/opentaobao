package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取ISV发起请求服务器IP APIResponse
taobao.appip.get

获取ISV发起请求服务器IP
*/
type TaobaoAppipGetAPIResponse struct {
    model.CommonResponse
    TaobaoAppipGetResponse
}

type TaobaoAppipGetResponse struct {
    XMLName xml.Name `xml:"appip_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ISV发起请求服务器IP
    
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`

    
}
