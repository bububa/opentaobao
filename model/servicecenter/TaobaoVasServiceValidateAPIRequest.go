package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasServiceValidateAPIRequest 增值服务订购服务验证 API请求
// taobao.vas.service.validate
//
// 增值服务订购服务验证
type TaobaoVasServiceValidateAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 服务编码
	_servCode string
}

// NewTaobaoVasServiceValidateRequest 初始化TaobaoVasServiceValidateAPIRequest对象
func NewTaobaoVasServiceValidateRequest() *TaobaoVasServiceValidateAPIRequest {
	return &TaobaoVasServiceValidateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVasServiceValidateAPIRequest) Reset() {
	r._nick = ""
	r._servCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasServiceValidateAPIRequest) GetApiMethodName() string {
	return "taobao.vas.service.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasServiceValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVasServiceValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoVasServiceValidateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoVasServiceValidateAPIRequest) GetNick() string {
	return r._nick
}

// SetServCode is ServCode Setter
// 服务编码
func (r *TaobaoVasServiceValidateAPIRequest) SetServCode(_servCode string) error {
	r._servCode = _servCode
	r.Set("serv_code", _servCode)
	return nil
}

// GetServCode ServCode Getter
func (r TaobaoVasServiceValidateAPIRequest) GetServCode() string {
	return r._servCode
}

var poolTaobaoVasServiceValidateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVasServiceValidateRequest()
	},
}

// GetTaobaoVasServiceValidateRequest 从 sync.Pool 获取 TaobaoVasServiceValidateAPIRequest
func GetTaobaoVasServiceValidateAPIRequest() *TaobaoVasServiceValidateAPIRequest {
	return poolTaobaoVasServiceValidateAPIRequest.Get().(*TaobaoVasServiceValidateAPIRequest)
}

// ReleaseTaobaoVasServiceValidateAPIRequest 将 TaobaoVasServiceValidateAPIRequest 放入 sync.Pool
func ReleaseTaobaoVasServiceValidateAPIRequest(v *TaobaoVasServiceValidateAPIRequest) {
	v.Reset()
	poolTaobaoVasServiceValidateAPIRequest.Put(v)
}
