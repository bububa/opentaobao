package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
申请免登Open Account Token API返回值 
taobao.open.account.token.apply

申请免登Open Account Token
*/
type TaobaoOpenAccountTokenApplyAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountTokenApplyResponse
}

// 申请免登Open Account Token 成功返回结果
type TaobaoOpenAccountTokenApplyResponse struct {
    XMLName xml.Name `xml:"open_account_token_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的token结果
    Data   *OpenAccountTokenApplyResult `json:"data,omitempty" xml:"data,omitempty"`
}
