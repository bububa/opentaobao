package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTicketStatusCallbackAPIRequest 代理商票状态查询回调 API请求
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
type TaobaoTrainAgentTicketStatusCallbackAPIRequest struct {
	model.Params
	// 描述信息
	_msg string
	// 唯一编号
	_id string
	// 车票状态
	_ticketStatus int64
}

// NewTaobaoTrainAgentTicketStatusCallbackRequest 初始化TaobaoTrainAgentTicketStatusCallbackAPIRequest对象
func NewTaobaoTrainAgentTicketStatusCallbackRequest() *TaobaoTrainAgentTicketStatusCallbackAPIRequest {
	return &TaobaoTrainAgentTicketStatusCallbackAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentTicketStatusCallbackAPIRequest) Reset() {
	r._msg = ""
	r._id = ""
	r._ticketStatus = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.ticket.status.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsg is Msg Setter
// 描述信息
func (r *TaobaoTrainAgentTicketStatusCallbackAPIRequest) SetMsg(_msg string) error {
	r._msg = _msg
	r.Set("msg", _msg)
	return nil
}

// GetMsg Msg Getter
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetMsg() string {
	return r._msg
}

// SetId is Id Setter
// 唯一编号
func (r *TaobaoTrainAgentTicketStatusCallbackAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetId() string {
	return r._id
}

// SetTicketStatus is TicketStatus Setter
// 车票状态
func (r *TaobaoTrainAgentTicketStatusCallbackAPIRequest) SetTicketStatus(_ticketStatus int64) error {
	r._ticketStatus = _ticketStatus
	r.Set("ticket_status", _ticketStatus)
	return nil
}

// GetTicketStatus TicketStatus Getter
func (r TaobaoTrainAgentTicketStatusCallbackAPIRequest) GetTicketStatus() int64 {
	return r._ticketStatus
}

var poolTaobaoTrainAgentTicketStatusCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentTicketStatusCallbackRequest()
	},
}

// GetTaobaoTrainAgentTicketStatusCallbackRequest 从 sync.Pool 获取 TaobaoTrainAgentTicketStatusCallbackAPIRequest
func GetTaobaoTrainAgentTicketStatusCallbackAPIRequest() *TaobaoTrainAgentTicketStatusCallbackAPIRequest {
	return poolTaobaoTrainAgentTicketStatusCallbackAPIRequest.Get().(*TaobaoTrainAgentTicketStatusCallbackAPIRequest)
}

// ReleaseTaobaoTrainAgentTicketStatusCallbackAPIRequest 将 TaobaoTrainAgentTicketStatusCallbackAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentTicketStatusCallbackAPIRequest(v *TaobaoTrainAgentTicketStatusCallbackAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentTicketStatusCallbackAPIRequest.Put(v)
}
