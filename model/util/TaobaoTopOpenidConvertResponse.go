package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
混淆nick转openid APIResponse
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
*/
type TaobaoTopOpenidConvertAPIResponse struct {
    model.CommonResponse
    TaobaoTopOpenidConvertResponse
}

type TaobaoTopOpenidConvertResponse struct {
    XMLName xml.Name `xml:"top_openid_convert_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // open_id
    
    OpenId   string `json:"open_id,omitempty" xml:"open_id,omitempty"`

    
}
