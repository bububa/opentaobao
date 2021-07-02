package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest 百川登录二次验证 API请求
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
type TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountLogindoublecheckRequest 初始化TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest对象
func NewTaobaoBaichuanOpenaccountLogindoublecheckRequest() *TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest {
	return &TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetName() string {
	return r._name
}
