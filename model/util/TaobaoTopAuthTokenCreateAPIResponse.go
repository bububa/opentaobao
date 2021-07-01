package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopAuthTokenCreateAPIResponse
获取Access Token API返回值
taobao.top.auth.token.create

用户通过code换获取access_token，https only */
type TaobaoTopAuthTokenCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTopAuthTokenCreateAPIResponseModel
}

// TaobaoTopAuthTokenCreateAPIResponseModel is 获取Access Token 成功返回结果
type TaobaoTopAuthTokenCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"top_auth_token_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致
	TokenResult string `json:"token_result,omitempty" xml:"token_result,omitempty"`
}
