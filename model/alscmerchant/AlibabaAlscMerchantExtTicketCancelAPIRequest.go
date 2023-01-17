package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketCancelAPIRequest 凭证作废 API请求
// alibaba.alsc.merchant.ext.ticket.cancel
//
// isv调用银行接口超时导致凭证信息同步失败,触发关单处理,调用作废接口
type AlibabaAlscMerchantExtTicketCancelAPIRequest struct {
	model.Params
	// 凭证任务单Id
	_ticketTaskId string
	// 请求的唯一requestId
	_ticketRequestId string
	// 购买商品的买家userId((纯数字))
	_buyerId string
}

// NewAlibabaAlscMerchantExtTicketCancelRequest 初始化AlibabaAlscMerchantExtTicketCancelAPIRequest对象
func NewAlibabaAlscMerchantExtTicketCancelRequest() *AlibabaAlscMerchantExtTicketCancelAPIRequest {
	return &AlibabaAlscMerchantExtTicketCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticket.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketTaskId is TicketTaskId Setter
// 凭证任务单Id
func (r *AlibabaAlscMerchantExtTicketCancelAPIRequest) SetTicketTaskId(_ticketTaskId string) error {
	r._ticketTaskId = _ticketTaskId
	r.Set("ticket_task_id", _ticketTaskId)
	return nil
}

// GetTicketTaskId TicketTaskId Getter
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetTicketTaskId() string {
	return r._ticketTaskId
}

// SetTicketRequestId is TicketRequestId Setter
// 请求的唯一requestId
func (r *AlibabaAlscMerchantExtTicketCancelAPIRequest) SetTicketRequestId(_ticketRequestId string) error {
	r._ticketRequestId = _ticketRequestId
	r.Set("ticket_request_id", _ticketRequestId)
	return nil
}

// GetTicketRequestId TicketRequestId Getter
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetTicketRequestId() string {
	return r._ticketRequestId
}

// SetBuyerId is BuyerId Setter
// 购买商品的买家userId((纯数字))
func (r *AlibabaAlscMerchantExtTicketCancelAPIRequest) SetBuyerId(_buyerId string) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r AlibabaAlscMerchantExtTicketCancelAPIRequest) GetBuyerId() string {
	return r._buyerId
}
