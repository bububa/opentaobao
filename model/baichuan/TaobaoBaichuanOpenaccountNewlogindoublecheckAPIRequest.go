package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest 百川新登录二次验证 API请求
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountNewlogindoublecheckRequest 初始化TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest对象
func NewTaobaoBaichuanOpenaccountNewlogindoublecheckRequest() *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest {
	return &TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.newlogindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest) GetName() string {
	return r._name
}
