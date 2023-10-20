package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarEticketConsumeAPIRequest 天猫汽车电子凭证核销 API请求
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
type TmallAliautoTradeCarEticketConsumeAPIRequest struct {
	model.Params
	// 入参
	_command *ConsumeEticketCommand
}

// NewTmallAliautoTradeCarEticketConsumeRequest 初始化TmallAliautoTradeCarEticketConsumeAPIRequest对象
func NewTmallAliautoTradeCarEticketConsumeRequest() *TmallAliautoTradeCarEticketConsumeAPIRequest {
	return &TmallAliautoTradeCarEticketConsumeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoTradeCarEticketConsumeAPIRequest) Reset() {
	r._command = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeCarEticketConsumeAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.eticket.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeCarEticketConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoTradeCarEticketConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *TmallAliautoTradeCarEticketConsumeAPIRequest) SetCommand(_command *ConsumeEticketCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallAliautoTradeCarEticketConsumeAPIRequest) GetCommand() *ConsumeEticketCommand {
	return r._command
}

var poolTmallAliautoTradeCarEticketConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoTradeCarEticketConsumeRequest()
	},
}

// GetTmallAliautoTradeCarEticketConsumeRequest 从 sync.Pool 获取 TmallAliautoTradeCarEticketConsumeAPIRequest
func GetTmallAliautoTradeCarEticketConsumeAPIRequest() *TmallAliautoTradeCarEticketConsumeAPIRequest {
	return poolTmallAliautoTradeCarEticketConsumeAPIRequest.Get().(*TmallAliautoTradeCarEticketConsumeAPIRequest)
}

// ReleaseTmallAliautoTradeCarEticketConsumeAPIRequest 将 TmallAliautoTradeCarEticketConsumeAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoTradeCarEticketConsumeAPIRequest(v *TmallAliautoTradeCarEticketConsumeAPIRequest) {
	v.Reset()
	poolTmallAliautoTradeCarEticketConsumeAPIRequest.Put(v)
}
