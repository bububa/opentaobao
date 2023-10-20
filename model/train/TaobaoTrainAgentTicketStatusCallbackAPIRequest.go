package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentticketstatuscallbackAPIRequest 代理商票状态查询回调 API请求
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
type TaobaotrainagentticketstatuscallbackAPIRequest struct {
	model.Params
	// 描述信息
	_msg string
	// 唯一编号
	_id string
	// 车票状态
	_ticketStatus int64
}

// NewTaobaotrainagentticketstatuscallbackRequest 初始化TaobaotrainagentticketstatuscallbackAPIRequest对象
func NewTaobaotrainagentticketstatuscallbackRequest() *TaobaotrainagentticketstatuscallbackAPIRequest {
	return &TaobaotrainagentticketstatuscallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.ticket.status.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsg is Msg Setter
// 描述信息
func (r *TaobaotrainagentticketstatuscallbackAPIRequest) SetMsg(_msg string) error {
	r._msg = _msg
	r.Set("msg", _msg)
	return nil
}

// GetMsg Msg Getter
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetMsg() string {
	return r._msg
}

// SetId is Id Setter
// 唯一编号
func (r *TaobaotrainagentticketstatuscallbackAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetId() string {
	return r._id
}

// SetTicketStatus is TicketStatus Setter
// 车票状态
func (r *TaobaotrainagentticketstatuscallbackAPIRequest) SetTicketStatus(_ticketStatus int64) error {
	r._ticketStatus = _ticketStatus
	r.Set("ticket_status", _ticketStatus)
	return nil
}

// GetTicketStatus TicketStatus Getter
func (r TaobaotrainagentticketstatuscallbackAPIRequest) GetTicketStatus() int64 {
	return r._ticketStatus
}
