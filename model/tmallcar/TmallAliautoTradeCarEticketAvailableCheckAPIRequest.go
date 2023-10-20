package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotradecareticketavailablecheckAPIRequest 天猫汽车电子凭证核销前校验 API请求
// tmall.aliauto.trade.car.eticket.available.check
//
// 天猫汽车核销码可用性校验
type TmallaliautotradecareticketavailablecheckAPIRequest struct {
	model.Params
	// 入参
	_command *CheckEticketAvailableCommand
}

// NewTmallaliautotradecareticketavailablecheckRequest 初始化TmallaliautotradecareticketavailablecheckAPIRequest对象
func NewTmallaliautotradecareticketavailablecheckRequest() *TmallaliautotradecareticketavailablecheckAPIRequest {
	return &TmallaliautotradecareticketavailablecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautotradecareticketavailablecheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.eticket.available.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautotradecareticketavailablecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautotradecareticketavailablecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *TmallaliautotradecareticketavailablecheckAPIRequest) SetCommand(_command *CheckEticketAvailableCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallaliautotradecareticketavailablecheckAPIRequest) GetCommand() *CheckEticketAvailableCommand {
	return r._command
}
