package tmallcar

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.restpayfee.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeRestpayfeeModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
