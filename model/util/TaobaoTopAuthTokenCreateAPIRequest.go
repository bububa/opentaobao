package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopauthtokencreateAPIRequest 获取Access Token API请求
// taobao.top.auth.token.create
//
// 用户通过code换获取access_token，https only
type TaobaotopauthtokencreateAPIRequest struct {
	model.Params
	// 授权code，grantType==authorization_code 时需要
	_code string
	// 与生成code的uuid配对
	_uuid string
}

// NewTaobaotopauthtokencreateRequest 初始化TaobaotopauthtokencreateAPIRequest对象
func NewTaobaotopauthtokencreateRequest() *TaobaotopauthtokencreateAPIRequest {
	return &TaobaotopauthtokencreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopauthtokencreateAPIRequest) GetApiMethodName() string {
	return "taobao.top.auth.token.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopauthtokencreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopauthtokencreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 授权code，grantType==authorization_code 时需要
func (r *TaobaotopauthtokencreateAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaotopauthtokencreateAPIRequest) GetCode() string {
	return r._code
}

// SetUuid is Uuid Setter
// 与生成code的uuid配对
func (r *TaobaotopauthtokencreateAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaotopauthtokencreateAPIRequest) GetUuid() string {
	return r._uuid
}
