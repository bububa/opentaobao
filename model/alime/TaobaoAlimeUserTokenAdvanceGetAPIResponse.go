package alime

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户免登录令牌v2 API返回值 
taobao.alime.user.token.advance.get

根据第三账号信息获取用户的免登录令牌
*/
type TaobaoAlimeUserTokenAdvanceGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlimeUserTokenAdvanceGetAPIResponseModel
}

// 获取用户免登录令牌v2 成功返回结果
type TaobaoAlimeUserTokenAdvanceGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alime_user_token_advance_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应消息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 响应数据
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 响应编码(由于"code"为top保留字，用code0以示区分，文档中均以code说明)，code == 0为成功，其它为失败
    Code0   int64 `json:"code0,omitempty" xml:"code0,omitempty"`
}
