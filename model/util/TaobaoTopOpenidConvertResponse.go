package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
混淆nick转openid APIResponse
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
*/
type TaobaoTopOpenidConvertAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTopOpenidConvertResponse `json:"taobao_top_openid_convert_response,omitempty"`
}

type TaobaoTopOpenidConvertResponse struct {

    // open_id
    OpenId   string `json:"open_id,omitempty"`

}
