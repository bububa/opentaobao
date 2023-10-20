package tmallcar

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoTradeCarEticketAvailableCheckAPIRequest) Reset() {
	r._command = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.eticket.available.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoTradeCarEticketAvailableCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallAliautoTradeCarEticketAvailableCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoTradeCarEticketAvailableCheckRequest()
	},
}

// GetTmallAliautoTradeCarEticketAvailableCheckRequest 从 sync.Pool 获取 TmallAliautoTradeCarEticketAvailableCheckAPIRequest
func GetTmallAliautoTradeCarEticketAvailableCheckAPIRequest() *TmallAliautoTradeCarEticketAvailableCheckAPIRequest {
	return poolTmallAliautoTradeCarEticketAvailableCheckAPIRequest.Get().(*TmallAliautoTradeCarEticketAvailableCheckAPIRequest)
}

// ReleaseTmallAliautoTradeCarEticketAvailableCheckAPIRequest 将 TmallAliautoTradeCarEticketAvailableCheckAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoTradeCarEticketAvailableCheckAPIRequest(v *TmallAliautoTradeCarEticketAvailableCheckAPIRequest) {
	v.Reset()
	poolTmallAliautoTradeCarEticketAvailableCheckAPIRequest.Put(v)
}
