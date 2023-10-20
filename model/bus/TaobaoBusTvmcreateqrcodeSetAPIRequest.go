package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmcreateqrcodeSetAPIRequest 自助机生成支付宝支付二维码 API请求
// taobao.bus.tvmcreateqrcode.set
//
// 用于汽车票线下自助机调用获取支付宝的二维码
type TaobaoBusTvmcreateqrcodeSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 超时时间（分钟）
	_timeoutExpress int64
}

// NewTaobaoBusTvmcreateqrcodeSetRequest 初始化TaobaoBusTvmcreateqrcodeSetAPIRequest对象
func NewTaobaoBusTvmcreateqrcodeSetRequest() *TaobaoBusTvmcreateqrcodeSetAPIRequest {
	return &TaobaoBusTvmcreateqrcodeSetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTvmcreateqrcodeSetAPIRequest) Reset() {
	r._alitripOrderId = ""
	r._timeoutExpress = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcreateqrcodeSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcreateqrcode.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcreateqrcodeSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTvmcreateqrcodeSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmcreateqrcodeSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaoBusTvmcreateqrcodeSetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

// SetTimeoutExpress is TimeoutExpress Setter
// 超时时间（分钟）
func (r *TaobaoBusTvmcreateqrcodeSetAPIRequest) SetTimeoutExpress(_timeoutExpress int64) error {
	r._timeoutExpress = _timeoutExpress
	r.Set("timeout_express", _timeoutExpress)
	return nil
}

// GetTimeoutExpress TimeoutExpress Getter
func (r TaobaoBusTvmcreateqrcodeSetAPIRequest) GetTimeoutExpress() int64 {
	return r._timeoutExpress
}

var poolTaobaoBusTvmcreateqrcodeSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTvmcreateqrcodeSetRequest()
	},
}

// GetTaobaoBusTvmcreateqrcodeSetRequest 从 sync.Pool 获取 TaobaoBusTvmcreateqrcodeSetAPIRequest
func GetTaobaoBusTvmcreateqrcodeSetAPIRequest() *TaobaoBusTvmcreateqrcodeSetAPIRequest {
	return poolTaobaoBusTvmcreateqrcodeSetAPIRequest.Get().(*TaobaoBusTvmcreateqrcodeSetAPIRequest)
}

// ReleaseTaobaoBusTvmcreateqrcodeSetAPIRequest 将 TaobaoBusTvmcreateqrcodeSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTvmcreateqrcodeSetAPIRequest(v *TaobaoBusTvmcreateqrcodeSetAPIRequest) {
	v.Reset()
	poolTaobaoBusTvmcreateqrcodeSetAPIRequest.Put(v)
}
