package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusrefundticketpricesetAPIRequest 汽车票退款申请接口 API请求
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
type TaobaobusrefundticketpricesetAPIRequest struct {
	model.Params
	// 退票申请入参
	_offlineRefundTicketRq *OfflineRefundTicketPriceRq
}

// NewTaobaobusrefundticketpricesetRequest 初始化TaobaobusrefundticketpricesetAPIRequest对象
func NewTaobaobusrefundticketpricesetRequest() *TaobaobusrefundticketpricesetAPIRequest {
	return &TaobaobusrefundticketpricesetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusrefundticketpricesetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refundticketprice.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusrefundticketpricesetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusrefundticketpricesetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineRefundTicketRq is OfflineRefundTicketRq Setter
// 退票申请入参
func (r *TaobaobusrefundticketpricesetAPIRequest) SetOfflineRefundTicketRq(_offlineRefundTicketRq *OfflineRefundTicketPriceRq) error {
	r._offlineRefundTicketRq = _offlineRefundTicketRq
	r.Set("offline_refund_ticket_rq", _offlineRefundTicketRq)
	return nil
}

// GetOfflineRefundTicketRq OfflineRefundTicketRq Getter
func (r TaobaobusrefundticketpricesetAPIRequest) GetOfflineRefundTicketRq() *OfflineRefundTicketPriceRq {
	return r._offlineRefundTicketRq
}
