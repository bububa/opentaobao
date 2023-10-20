package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountloginbytokenAPIRequest 百川TOKEN 登录 API请求
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
type TaobaobaichuanopenaccountloginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanopenaccountloginbytokenRequest 初始化TaobaobaichuanopenaccountloginbytokenAPIRequest对象
func NewTaobaobaichuanopenaccountloginbytokenRequest() *TaobaobaichuanopenaccountloginbytokenAPIRequest {
	return &TaobaobaichuanopenaccountloginbytokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanopenaccountloginbytokenAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.loginbytoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanopenaccountloginbytokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanopenaccountloginbytokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanopenaccountloginbytokenAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanopenaccountloginbytokenAPIRequest) GetName() string {
	return r._name
}
