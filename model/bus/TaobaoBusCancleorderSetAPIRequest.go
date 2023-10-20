package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusCancleorderSetAPIRequest 取消订单 API请求
// taobao.bus.cancleorder.set
//
// 取消订单
type TaobaoBusCancleorderSetAPIRequest struct {
	model.Params
	// 阿里订单号
	_aliOrderId string
}

// NewTaobaoBusCancleorderSetRequest 初始化TaobaoBusCancleorderSetAPIRequest对象
func NewTaobaoBusCancleorderSetRequest() *TaobaoBusCancleorderSetAPIRequest {
	return &TaobaoBusCancleorderSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusCancleorderSetAPIRequest) Reset() {
	r._aliOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusCancleorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.cancleorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusCancleorderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusCancleorderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliOrderId is AliOrderId Setter
// 阿里订单号
func (r *TaobaoBusCancleorderSetAPIRequest) SetAliOrderId(_aliOrderId string) error {
	r._aliOrderId = _aliOrderId
	r.Set("ali_order_id", _aliOrderId)
	return nil
}

// GetAliOrderId AliOrderId Getter
func (r TaobaoBusCancleorderSetAPIRequest) GetAliOrderId() string {
	return r._aliOrderId
}

var poolTaobaoBusCancleorderSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusCancleorderSetRequest()
	},
}

// GetTaobaoBusCancleorderSetRequest 从 sync.Pool 获取 TaobaoBusCancleorderSetAPIRequest
func GetTaobaoBusCancleorderSetAPIRequest() *TaobaoBusCancleorderSetAPIRequest {
	return poolTaobaoBusCancleorderSetAPIRequest.Get().(*TaobaoBusCancleorderSetAPIRequest)
}

// ReleaseTaobaoBusCancleorderSetAPIRequest 将 TaobaoBusCancleorderSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusCancleorderSetAPIRequest(v *TaobaoBusCancleorderSetAPIRequest) {
	v.Reset()
	poolTaobaoBusCancleorderSetAPIRequest.Put(v)
}
