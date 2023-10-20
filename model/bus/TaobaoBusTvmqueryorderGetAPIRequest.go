package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmqueryorderGetAPIRequest 线下自助机查询订单信息 API请求
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
type TaobaoBusTvmqueryorderGetAPIRequest struct {
	model.Params
	// 阿里订单标编号
	_alitripOrderId string
}

// NewTaobaoBusTvmqueryorderGetRequest 初始化TaobaoBusTvmqueryorderGetAPIRequest对象
func NewTaobaoBusTvmqueryorderGetRequest() *TaobaoBusTvmqueryorderGetAPIRequest {
	return &TaobaoBusTvmqueryorderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTvmqueryorderGetAPIRequest) Reset() {
	r._alitripOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmqueryorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 阿里订单标编号
func (r *TaobaoBusTvmqueryorderGetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

var poolTaobaoBusTvmqueryorderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTvmqueryorderGetRequest()
	},
}

// GetTaobaoBusTvmqueryorderGetRequest 从 sync.Pool 获取 TaobaoBusTvmqueryorderGetAPIRequest
func GetTaobaoBusTvmqueryorderGetAPIRequest() *TaobaoBusTvmqueryorderGetAPIRequest {
	return poolTaobaoBusTvmqueryorderGetAPIRequest.Get().(*TaobaoBusTvmqueryorderGetAPIRequest)
}

// ReleaseTaobaoBusTvmqueryorderGetAPIRequest 将 TaobaoBusTvmqueryorderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTvmqueryorderGetAPIRequest(v *TaobaoBusTvmqueryorderGetAPIRequest) {
	v.Reset()
	poolTaobaoBusTvmqueryorderGetAPIRequest.Put(v)
}
