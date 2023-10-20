package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLoginbytokenAPIRequest 百川TOKEN 登录 API请求
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
type TaobaoBaichuanOpenaccountLoginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountLoginbytokenRequest 初始化TaobaoBaichuanOpenaccountLoginbytokenAPIRequest对象
func NewTaobaoBaichuanOpenaccountLoginbytokenRequest() *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest {
	return &TaobaoBaichuanOpenaccountLoginbytokenAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) Reset() {
	r._name = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.loginbytoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetName() string {
	return r._name
}

var poolTaobaoBaichuanOpenaccountLoginbytokenAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaichuanOpenaccountLoginbytokenRequest()
	},
}

// GetTaobaoBaichuanOpenaccountLoginbytokenRequest 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLoginbytokenAPIRequest
func GetTaobaoBaichuanOpenaccountLoginbytokenAPIRequest() *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest {
	return poolTaobaoBaichuanOpenaccountLoginbytokenAPIRequest.Get().(*TaobaoBaichuanOpenaccountLoginbytokenAPIRequest)
}

// ReleaseTaobaoBaichuanOpenaccountLoginbytokenAPIRequest 将 TaobaoBaichuanOpenaccountLoginbytokenAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLoginbytokenAPIRequest(v *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLoginbytokenAPIRequest.Put(v)
}
