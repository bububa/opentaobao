package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmcancelorderSetAPIRequest 线下自助机未付款取消订单 API请求
// taobao.bus.tvmcancelorder.set
//
// 自助机汽车票未付款取消订单
type TaobaoBusTvmcancelorderSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaoBusTvmcancelorderSetRequest 初始化TaobaoBusTvmcancelorderSetAPIRequest对象
func NewTaobaoBusTvmcancelorderSetRequest() *TaobaoBusTvmcancelorderSetAPIRequest {
	return &TaobaoBusTvmcancelorderSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTvmcancelorderSetAPIRequest) Reset() {
	r._alitripOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcancelorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcancelorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcancelorderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTvmcancelorderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmcancelorderSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaoBusTvmcancelorderSetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

var poolTaobaoBusTvmcancelorderSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTvmcancelorderSetRequest()
	},
}

// GetTaobaoBusTvmcancelorderSetRequest 从 sync.Pool 获取 TaobaoBusTvmcancelorderSetAPIRequest
func GetTaobaoBusTvmcancelorderSetAPIRequest() *TaobaoBusTvmcancelorderSetAPIRequest {
	return poolTaobaoBusTvmcancelorderSetAPIRequest.Get().(*TaobaoBusTvmcancelorderSetAPIRequest)
}

// ReleaseTaobaoBusTvmcancelorderSetAPIRequest 将 TaobaoBusTvmcancelorderSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTvmcancelorderSetAPIRequest(v *TaobaoBusTvmcancelorderSetAPIRequest) {
	v.Reset()
	poolTaobaoBusTvmcancelorderSetAPIRequest.Put(v)
}
