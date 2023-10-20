package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountTokenValidateAPIRequest open account token验证 API请求
// taobao.open.account.token.validate
//
// open account token验证
type TaobaoOpenAccountTokenValidateAPIRequest struct {
	model.Params
	// token
	_paramToken string
}

// NewTaobaoOpenAccountTokenValidateRequest 初始化TaobaoOpenAccountTokenValidateAPIRequest对象
func NewTaobaoOpenAccountTokenValidateRequest() *TaobaoOpenAccountTokenValidateAPIRequest {
	return &TaobaoOpenAccountTokenValidateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenAccountTokenValidateAPIRequest) Reset() {
	r._paramToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.token.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamToken is ParamToken Setter
// token
func (r *TaobaoOpenAccountTokenValidateAPIRequest) SetParamToken(_paramToken string) error {
	r._paramToken = _paramToken
	r.Set("param_token", _paramToken)
	return nil
}

// GetParamToken ParamToken Getter
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetParamToken() string {
	return r._paramToken
}

var poolTaobaoOpenAccountTokenValidateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenAccountTokenValidateRequest()
	},
}

// GetTaobaoOpenAccountTokenValidateRequest 从 sync.Pool 获取 TaobaoOpenAccountTokenValidateAPIRequest
func GetTaobaoOpenAccountTokenValidateAPIRequest() *TaobaoOpenAccountTokenValidateAPIRequest {
	return poolTaobaoOpenAccountTokenValidateAPIRequest.Get().(*TaobaoOpenAccountTokenValidateAPIRequest)
}

// ReleaseTaobaoOpenAccountTokenValidateAPIRequest 将 TaobaoOpenAccountTokenValidateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenAccountTokenValidateAPIRequest(v *TaobaoOpenAccountTokenValidateAPIRequest) {
	v.Reset()
	poolTaobaoOpenAccountTokenValidateAPIRequest.Put(v)
}
