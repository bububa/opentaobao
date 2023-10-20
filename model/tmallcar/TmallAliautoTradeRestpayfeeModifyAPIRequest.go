package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeRestpayfeeModifyAPIRequest 分阶段订单尾款改价 API请求
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
type TmallAliautoTradeRestpayfeeModifyAPIRequest struct {
	model.Params
	// 入参
	_command *ModifyRestPaymentCommand
}

// NewTmallAliautoTradeRestpayfeeModifyRequest 初始化TmallAliautoTradeRestpayfeeModifyAPIRequest对象
func NewTmallAliautoTradeRestpayfeeModifyRequest() *TmallAliautoTradeRestpayfeeModifyAPIRequest {
	return &TmallAliautoTradeRestpayfeeModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoTradeRestpayfeeModifyAPIRequest) Reset() {
	r._command = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.restpayfee.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *TmallAliautoTradeRestpayfeeModifyAPIRequest) SetCommand(_command *ModifyRestPaymentCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetCommand() *ModifyRestPaymentCommand {
	return r._command
}

var poolTmallAliautoTradeRestpayfeeModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoTradeRestpayfeeModifyRequest()
	},
}

// GetTmallAliautoTradeRestpayfeeModifyRequest 从 sync.Pool 获取 TmallAliautoTradeRestpayfeeModifyAPIRequest
func GetTmallAliautoTradeRestpayfeeModifyAPIRequest() *TmallAliautoTradeRestpayfeeModifyAPIRequest {
	return poolTmallAliautoTradeRestpayfeeModifyAPIRequest.Get().(*TmallAliautoTradeRestpayfeeModifyAPIRequest)
}

// ReleaseTmallAliautoTradeRestpayfeeModifyAPIRequest 将 TmallAliautoTradeRestpayfeeModifyAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoTradeRestpayfeeModifyAPIRequest(v *TmallAliautoTradeRestpayfeeModifyAPIRequest) {
	v.Reset()
	poolTmallAliautoTradeRestpayfeeModifyAPIRequest.Put(v)
}
