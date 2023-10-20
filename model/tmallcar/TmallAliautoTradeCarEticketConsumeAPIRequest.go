package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotradecareticketconsumeAPIRequest 天猫汽车电子凭证核销 API请求
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
type TmallaliautotradecareticketconsumeAPIRequest struct {
	model.Params
	// 入参
	_command *ConsumeEticketCommand
}

// NewTmallaliautotradecareticketconsumeRequest 初始化TmallaliautotradecareticketconsumeAPIRequest对象
func NewTmallaliautotradecareticketconsumeRequest() *TmallaliautotradecareticketconsumeAPIRequest {
	return &TmallaliautotradecareticketconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautotradecareticketconsumeAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.eticket.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautotradecareticketconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautotradecareticketconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *TmallaliautotradecareticketconsumeAPIRequest) SetCommand(_command *ConsumeEticketCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallaliautotradecareticketconsumeAPIRequest) GetCommand() *ConsumeEticketCommand {
	return r._command
}
