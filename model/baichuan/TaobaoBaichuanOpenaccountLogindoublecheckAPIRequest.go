package baichuan

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.logindoublecheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountLogindoublecheckRequest()
	},
}

// GetTaobaoBaichuanOpenaccountLogindoublecheckRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest
func GetTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest() *TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest {
	return poolTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest.Get().(*TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest 将 TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest(v *TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLogindoublecheckAPIRequest.Put(v)
}
