package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
open account token验证 APIResponse
taobao.open.account.token.validate

open account token验证
*/
type TaobaoOpenAccountTokenValidateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountTokenValidateResponse
}

type TaobaoOpenAccountTokenValidateResponse struct {
    XMLName xml.Name `xml:"open_account_token_validate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 验证成功返回token中的信息
    
    Data   *OpenAccountTokenValidateResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
