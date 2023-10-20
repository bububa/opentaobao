package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeDeliverAPIRequest 飞猪度假-订单发货接口 API请求
// alitrip.travel.trade.deliver
//
// 航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
type AlitripTravelTradeDeliverAPIRequest struct {
	model.Params
	// 子订单id
	_subOrderId int64
}

// NewAlitripTravelTradeDeliverRequest 初始化AlitripTravelTradeDeliverAPIRequest对象
func NewAlitripTravelTradeDeliverRequest() *AlitripTravelTradeDeliverAPIRequest {
	return &AlitripTravelTradeDeliverAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelTradeDeliverAPIRequest) Reset() {
	r._subOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeDeliverAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.deliver"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeDeliverAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelTradeDeliverAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderId is SubOrderId Setter
// 子订单id
func (r *AlitripTravelTradeDeliverAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlitripTravelTradeDeliverAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

var poolAlitripTravelTradeDeliverAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelTradeDeliverRequest()
	},
}

// GetAlitripTravelTradeDeliverRequest 从 sync.Pool 获取 AlitripTravelTradeDeliverAPIRequest
func GetAlitripTravelTradeDeliverAPIRequest() *AlitripTravelTradeDeliverAPIRequest {
	return poolAlitripTravelTradeDeliverAPIRequest.Get().(*AlitripTravelTradeDeliverAPIRequest)
}

// ReleaseAlitripTravelTradeDeliverAPIRequest 将 AlitripTravelTradeDeliverAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelTradeDeliverAPIRequest(v *AlitripTravelTradeDeliverAPIRequest) {
	v.Reset()
	poolAlitripTravelTradeDeliverAPIRequest.Put(v)
}
