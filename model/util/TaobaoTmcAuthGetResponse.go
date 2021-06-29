package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
TMC授权token API返回值 
taobao.tmc.auth.get

TMC连接授权Token
*/
type TaobaoTmcAuthGetAPIResponse struct {
    model.CommonResponse
    TaobaoTmcAuthGetResponse
}

// TMC授权token 成功返回结果
type TaobaoTmcAuthGetResponse struct {
    XMLName xml.Name `xml:"tmc_auth_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
