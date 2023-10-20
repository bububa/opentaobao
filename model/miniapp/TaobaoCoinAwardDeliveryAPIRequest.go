package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCoinAwardDeliveryAPIRequest 淘金币奖励投放 API请求
// taobao.coin.award.delivery
//
// 淘金币奖励投放
type TaobaoCoinAwardDeliveryAPIRequest struct {
	model.Params
	// openid
	_param0 string
	// 渠道ID
	_param1 int64
	// 订单信息
	_orderDto *OrderDto
}

// NewTaobaoCoinAwardDeliveryRequest 初始化TaobaoCoinAwardDeliveryAPIRequest对象
func NewTaobaoCoinAwardDeliveryRequest() *TaobaoCoinAwardDeliveryAPIRequest {
	return &TaobaoCoinAwardDeliveryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCoinAwardDeliveryAPIRequest) Reset() {
	r._param0 = ""
	r._param1 = 0
	r._orderDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCoinAwardDeliveryAPIRequest) GetApiMethodName() string {
	return "taobao.coin.award.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCoinAwardDeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCoinAwardDeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// openid
func (r *TaobaoCoinAwardDeliveryAPIRequest) SetParam0(_param0 string) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoCoinAwardDeliveryAPIRequest) GetParam0() string {
	return r._param0
}

// SetParam1 is Param1 Setter
// 渠道ID
func (r *TaobaoCoinAwardDeliveryAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoCoinAwardDeliveryAPIRequest) GetParam1() int64 {
	return r._param1
}

// SetOrderDto is OrderDto Setter
// 订单信息
func (r *TaobaoCoinAwardDeliveryAPIRequest) SetOrderDto(_orderDto *OrderDto) error {
	r._orderDto = _orderDto
	r.Set("order_dto", _orderDto)
	return nil
}

// GetOrderDto OrderDto Getter
func (r TaobaoCoinAwardDeliveryAPIRequest) GetOrderDto() *OrderDto {
	return r._orderDto
}

var poolTaobaoCoinAwardDeliveryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCoinAwardDeliveryRequest()
	},
}

// GetTaobaoCoinAwardDeliveryRequest 从 sync.Pool 获取 TaobaoCoinAwardDeliveryAPIRequest
func GetTaobaoCoinAwardDeliveryAPIRequest() *TaobaoCoinAwardDeliveryAPIRequest {
	return poolTaobaoCoinAwardDeliveryAPIRequest.Get().(*TaobaoCoinAwardDeliveryAPIRequest)
}

// ReleaseTaobaoCoinAwardDeliveryAPIRequest 将 TaobaoCoinAwardDeliveryAPIRequest 放入 sync.Pool
func ReleaseTaobaoCoinAwardDeliveryAPIRequest(v *TaobaoCoinAwardDeliveryAPIRequest) {
	v.Reset()
	poolTaobaoCoinAwardDeliveryAPIRequest.Put(v)
}
