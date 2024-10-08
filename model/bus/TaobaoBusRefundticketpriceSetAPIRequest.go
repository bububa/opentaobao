package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundticketpriceSetAPIRequest 汽车票退款申请接口 API请求
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
type TaobaoBusRefundticketpriceSetAPIRequest struct {
	model.Params
	// 退票申请入参
	_offlineRefundTicketRq *OfflineRefundTicketPriceRq
}

// NewTaobaoBusRefundticketpriceSetRequest 初始化TaobaoBusRefundticketpriceSetAPIRequest对象
func NewTaobaoBusRefundticketpriceSetRequest() *TaobaoBusRefundticketpriceSetAPIRequest {
	return &TaobaoBusRefundticketpriceSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusRefundticketpriceSetAPIRequest) Reset() {
	r._offlineRefundTicketRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refundticketprice.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineRefundTicketRq is OfflineRefundTicketRq Setter
// 退票申请入参
func (r *TaobaoBusRefundticketpriceSetAPIRequest) SetOfflineRefundTicketRq(_offlineRefundTicketRq *OfflineRefundTicketPriceRq) error {
	r._offlineRefundTicketRq = _offlineRefundTicketRq
	r.Set("offline_refund_ticket_rq", _offlineRefundTicketRq)
	return nil
}

// GetOfflineRefundTicketRq OfflineRefundTicketRq Getter
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetOfflineRefundTicketRq() *OfflineRefundTicketPriceRq {
	return r._offlineRefundTicketRq
}

var poolTaobaoBusRefundticketpriceSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusRefundticketpriceSetRequest()
	},
}

// GetTaobaoBusRefundticketpriceSetRequest 从 sync.Pool 获取 TaobaoBusRefundticketpriceSetAPIRequest
func GetTaobaoBusRefundticketpriceSetAPIRequest() *TaobaoBusRefundticketpriceSetAPIRequest {
	return poolTaobaoBusRefundticketpriceSetAPIRequest.Get().(*TaobaoBusRefundticketpriceSetAPIRequest)
}

// ReleaseTaobaoBusRefundticketpriceSetAPIRequest 将 TaobaoBusRefundticketpriceSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusRefundticketpriceSetAPIRequest(v *TaobaoBusRefundticketpriceSetAPIRequest) {
	v.Reset()
	poolTaobaoBusRefundticketpriceSetAPIRequest.Put(v)
}
