package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarEticketAvailableCheckAPIRequest 天猫汽车电子凭证核销前校验 API请求
// tmall.aliauto.trade.car.eticket.available.check
//
// 天猫汽车核销码可用性校验
type TmallAliautoTradeCarEticketAvailableCheckAPIRequest struct {
	model.Params
	// 入参
	_command *CheckEticketAvailableCommand
}

// NewTmallAliautoTradeCarEticketAvailableCheckRequest 初始化TmallAliautoTradeCarEticketAvailableCheckAPIRequest对象
func NewTmallAliautoTradeCarEticketAvailableCheckRequest() *TmallAliautoTradeCarEticketAvailableCheckAPIRequest {
	return &TmallAliautoTradeCarEticketAvailableCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.eticket.available.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCommand is Command Setter
// 入参
func (r *TmallAliautoTradeCarEticketAvailableCheckAPIRequest) SetCommand(_command *CheckEticketAvailableCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetCommand() *CheckEticketAvailableCommand {
	return r._command
}
