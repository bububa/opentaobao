package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanUserLogindoublecheckAPIRequest 百川H5登录二次验证 API请求
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
type TaobaoBaichuanUserLogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanUserLogindoublecheckRequest 初始化TaobaoBaichuanUserLogindoublecheckAPIRequest对象
func NewTaobaoBaichuanUserLogindoublecheckRequest() *TaobaoBaichuanUserLogindoublecheckAPIRequest {
	return &TaobaoBaichuanUserLogindoublecheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanUserLogindoublecheckAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.user.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanUserLogindoublecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanUserLogindoublecheckAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanUserLogindoublecheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanUserLogindoublecheckRequest()
	},
}

// GetTaobaoBaichuanUserLogindoublecheckRequest 从 sync.Pool 获取 TaobaoBaichuanUserLogindoublecheckAPIRequest
func GetTaobaoBaichuanUserLogindoublecheckAPIRequest() *TaobaoBaichuanUserLogindoublecheckAPIRequest {
	return poolTaobaoBaichuanUserLogindoublecheckAPIRequest.Get().(*TaobaoBaichuanUserLogindoublecheckAPIRequest)
}

// ReleaseTaobaoBaichuanUserLogindoublecheckAPIRequest 将 TaobaoBaichuanUserLogindoublecheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanUserLogindoublecheckAPIRequest(v *TaobaoBaichuanUserLogindoublecheckAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanUserLogindoublecheckAPIRequest.Put(v)
}
