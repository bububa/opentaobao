package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
申请免登Open Account Token APIResponse
taobao.open.account.token.apply

申请免登Open Account Token
*/
type TaobaoOpenAccountTokenApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"open_account_token_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的token结果
    
    Data   *OpenAccountTokenApplyResult `json:"data,omitempty" xml:"