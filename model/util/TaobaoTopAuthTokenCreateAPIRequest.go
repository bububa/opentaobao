package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopAuthTokenCreateAPIRequest 获取Access Token API请求
// taobao.top.auth.token.create
//
// 用户通过code换获取access_token，https only
type TaobaoTopAuthTokenCreateAPIRequest struct {
	model.Params
	// 授权code，grantType==authorization_code 时需要
	_code string
	// 与生成code的uuid配对
	_uuid string
}

// NewTaobaoTopAuthTokenCreateRequest 初始化TaobaoTopAuthTokenCreateAPIRequest对象
func NewTaobaoTopAuthTokenCreateRequest() *TaobaoTopAuthTokenCreateAPIRequest {
	return &TaobaoTopAuthTokenCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopAuthTokenCreateAPIRequest) GetApiMethodName() string {
	return "taobao.top.auth.token.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopAuthTokenCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCode is Code Setter
// 授权code，grantType==authorization_code 时需要
func (r *TaobaoTopAuthTokenCreateAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoTopAuthTokenCreateAPIRequest) GetCode() string {
	return r._code
}

// SetUuid is Uuid Setter
// 与生成code的uuid配对
func (r *TaobaoTopAuthTokenCreateAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoTopAuthTokenCreateAPIRequest) GetUuid() string {
	return r._uuid
}
