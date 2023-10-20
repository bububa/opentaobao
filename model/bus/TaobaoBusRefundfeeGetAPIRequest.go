package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundfeeGetAPIRequest 查询退票费用明细 API请求
// taobao.bus.refundfee.get
//
// 查询退票的费用信息
type TaobaoBusRefundfeeGetAPIRequest struct {
	model.Params
	// 飞猪子订单号
	_subOrderIds []int64
	// 飞猪订单号
	_aliTripOrderId string
}

// NewTaobaoBusRefundfeeGetRequest 初始化TaobaoBusRefundfeeGetAPIRequest对象
func NewTaobaoBusRefundfeeGetRequest() *TaobaoBusRefundfeeGetAPIRequest {
	return &TaobaoBusRefundfeeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusRefundfeeGetAPIRequest) Reset() {
	r._subOrderIds = r._subOrderIds[:0]
	r._aliTripOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundfeeGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refundfee.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundfeeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusRefundfeeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderIds is SubOrderIds Setter
// 飞猪子订单号
func (r *TaobaoBusRefundfeeGetAPIRequest) SetSubOrderIds(_subOrderIds []int64) error {
	r._subOrderIds = _subOrderIds
	r.Set("sub_order_ids", _subOrderIds)
	return nil
}

// GetSubOrderIds SubOrderIds Getter
func (r TaobaoBusRefundfeeGetAPIRequest) GetSubOrderIds() []int64 {
	return r._subOrderIds
}

// SetAliTripOrderId is AliTripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusRefundfeeGetAPIRequest) SetAliTripOrderId(_aliTripOrderId string) error {
	r._aliTripOrderId = _aliTripOrderId
	r.Set("ali_trip_order_id", _aliTripOrderId)
	return nil
}

// GetAliTripOrderId AliTripOrderId Getter
func (r TaobaoBusRefundfeeGetAPIRequest) GetAliTripOrderId() string {
	return r._aliTripOrderId
}

var poolTaobaoBusRefundfeeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusRefundfeeGetRequest()
	},
}

// GetTaobaoBusRefundfeeGetRequest 从 sync.Pool 获取 TaobaoBusRefundfeeGetAPIRequest
func GetTaobaoBusRefundfeeGetAPIRequest() *TaobaoBusRefundfeeGetAPIRequest {
	return poolTaobaoBusRefundfeeGetAPIRequest.Get().(*TaobaoBusRefundfeeGetAPIRequest)
}

// ReleaseTaobaoBusRefundfeeGetAPIRequest 将 TaobaoBusRefundfeeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusRefundfeeGetAPIRequest(v *TaobaoBusRefundfeeGetAPIRequest) {
	v.Reset()
	poolTaobaoBusRefundfeeGetAPIRequest.Put(v)
}
