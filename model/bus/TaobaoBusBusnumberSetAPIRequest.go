package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusBusnumberSetAPIRequest 商家汽车票车次更新通知接口 API请求
// taobao.bus.busnumber.set
//
// 商家汽车票车次更新后，调用该接口通知平台。
type TaobaoBusBusnumberSetAPIRequest struct {
	model.Params
	// 车次更新通知参数
	_pushParam *TopBusNumerPushRq
}

// NewTaobaoBusBusnumberSetRequest 初始化TaobaoBusBusnumberSetAPIRequest对象
func NewTaobaoBusBusnumberSetRequest() *TaobaoBusBusnumberSetAPIRequest {
	return &TaobaoBusBusnumberSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusBusnumberSetAPIRequest) Reset() {
	r._pushParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusBusnumberSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushParam is PushParam Setter
// 车次更新通知参数
func (r *TaobaoBusBusnumberSetAPIRequest) SetPushParam(_pushParam *TopBusNumerPushRq) error {
	r._pushParam = _pushParam
	r.Set("push_param", _pushParam)
	return nil
}

// GetPushParam PushParam Getter
func (r TaobaoBusBusnumberSetAPIRequest) GetPushParam() *TopBusNumerPushRq {
	return r._pushParam
}

var poolTaobaoBusBusnumberSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusBusnumberSetRequest()
	},
}

// GetTaobaoBusBusnumberSetRequest 从 sync.Pool 获取 TaobaoBusBusnumberSetAPIRequest
func GetTaobaoBusBusnumberSetAPIRequest() *TaobaoBusBusnumberSetAPIRequest {
	return poolTaobaoBusBusnumberSetAPIRequest.Get().(*TaobaoBusBusnumberSetAPIRequest)
}

// ReleaseTaobaoBusBusnumberSetAPIRequest 将 TaobaoBusBusnumberSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusBusnumberSetAPIRequest(v *TaobaoBusBusnumberSetAPIRequest) {
	v.Reset()
	poolTaobaoBusBusnumberSetAPIRequest.Put(v)
}
