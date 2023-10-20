package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotraderestpayfeemodifyAPIRequest 分阶段订单尾款改价 API请求
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
type TmallaliautotraderestpayfeemodifyAPIRequest struct {
	model.Params
	// 入参
	_command *ModifyRestPaymentCommand
}

// NewTmallaliautotraderestpayfeemodifyRequest 初始化TmallaliautotraderestpayfeemodifyAPIRequest对象
func NewTmallaliautotraderestpayfeemodifyRequest() *TmallaliautotraderestpayfeemodifyAPIRequest {
	return &TmallaliautotraderestpayfeemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautotraderestpayfeemodifyAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.restpayfee.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautotraderestpayfeemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautotraderestpayfeemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *TmallaliautotraderestpayfeemodifyAPIRequest) SetCommand(_command *ModifyRestPaymentCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallaliautotraderestpayfeemodifyAPIRequest) GetCommand() *ModifyRestPaymentCommand {
	return r._command
}
