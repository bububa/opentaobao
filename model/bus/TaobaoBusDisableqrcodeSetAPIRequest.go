package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusDisableqrcodeSetAPIRequest 自助机失效二维码 API请求
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
type TaobaoBusDisableqrcodeSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaoBusDisableqrcodeSetRequest 初始化TaobaoBusDisableqrcodeSetAPIRequest对象
func NewTaobaoBusDisableqrcodeSetRequest() *TaobaoBusDisableqrcodeSetAPIRequest {
	return &TaobaoBusDisableqrcodeSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusDisableqrcodeSetAPIRequest) Reset() {
	r._alitripOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.disableqrcode.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusDisableqrcodeSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

var poolTaobaoBusDisableqrcodeSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusDisableqrcodeSetRequest()
	},
}

// GetTaobaoBusDisableqrcodeSetRequest 从 sync.Pool 获取 TaobaoBusDisableqrcodeSetAPIRequest
func GetTaobaoBusDisableqrcodeSetAPIRequest() *TaobaoBusDisableqrcodeSetAPIRequest {
	return poolTaobaoBusDisableqrcodeSetAPIRequest.Get().(*TaobaoBusDisableqrcodeSetAPIRequest)
}

// ReleaseTaobaoBusDisableqrcodeSetAPIRequest 将 TaobaoBusDisableqrcodeSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusDisableqrcodeSetAPIRequest(v *TaobaoBusDisableqrcodeSetAPIRequest) {
	v.Reset()
	poolTaobaoBusDisableqrcodeSetAPIRequest.Put(v)
}
