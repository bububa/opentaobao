package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建群 APIResponse
taobao.openim.tribe.create

创建一个openim的群
*/
type TaobaoOpenimTribeCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeCreateResponse
}

type TaobaoOpenimTribeCreateResponse struct {
    XMLName xml.Name `xml:"openim_tribe_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建群的信息
    
    TribeInfo   *TribeInfo `json:"tribe_info,omitempty" xml:"tribe_info,omitempty"`

    
}
