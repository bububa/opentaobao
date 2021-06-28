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
	RequestId     string         `json:"request_id,omitempty" xml:"top_openid_convert_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // open_id
    
    OpenId   string `json:"open_id,omitempty" xml:"