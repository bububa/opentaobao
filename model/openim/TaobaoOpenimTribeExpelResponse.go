package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群踢出成员 API返回值 
taobao.openim.tribe.expel

OPENIM群踢出成员
*/
type TaobaoOpenimTribeExpelAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeExpelResponse
}

// OPENIM群踢出成员 成功返回结果
type TaobaoOpenimTribeExpelResponse struct {
    XMLName xml.Name `xml:"openim_tribe_expel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
