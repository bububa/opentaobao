package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取Access Token API请求
taobao.top.auth.token.create

用户通过code换获取access_token，https only
*/
type TaobaoTopAuthTokenCreateRequest struct {
    model.Params
    // 授权code，grantType==authorization_code 时需要
    code   string
    // 与生成code的uuid配对
    uuid   string
}

// 初始化TaobaoTopAuthTokenCreateRequest对象
func NewTaobaoTopAuthTokenCreateRequest() *TaobaoTopAuthTokenCreateRequest{
    return &TaobaoTopAuthTokenCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopAuthTokenCreateRequest) GetApiMethodName() string {
    return "taobao.top.auth.token.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopAuthTokenCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 授权code，grantType==authorization_code 时需要
func (r *TaobaoTopAuthTokenCreateRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoTopAuthTokenCreateRequest) GetCode() string {
    return r.code
}
// Uuid Setter
// 与生成code的uuid配对
func (r *TaobaoTopAuthTokenCreateRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoTopAuthTokenCreateRequest) GetUuid() string {
    return r.uuid
}
