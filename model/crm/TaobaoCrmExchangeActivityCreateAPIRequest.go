package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeActivityCreateAPIRequest 创建积分兑换活动 API请求
// taobao.crm.exchange.activity.create
//
// 创建针对积分兑换类型的活动
type TaobaoCrmExchangeActivityCreateAPIRequest struct {
	model.Params
	// 创建积分兑换活动
	_exchangeActivityCreateDto *ExchangeActivityCreateDto
}

// NewTaobaoCrmExchangeActivityCreateRequest 初始化TaobaoCrmExchangeActivityCreateAPIRequest对象
func NewTaobaoCrmExchangeActivityCreateRequest() *TaobaoCrmExchangeActivityCreateAPIRequest {
	return &TaobaoCrmExchangeActivityCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmExchangeActivityCreateAPIRequest) Reset() {
	r._exchangeActivityCreateDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExchangeActivityCreateDto is ExchangeActivityCreateDto Setter
// 创建积分兑换活动
func (r *TaobaoCrmExchangeActivityCreateAPIRequest) SetExchangeActivityCreateDto(_exchangeActivityCreateDto *ExchangeActivityCreateDto) error {
	r._exchangeActivityCreateDto = _exchangeActivityCreateDto
	r.Set("exchange_activity_create_dto", _exchangeActivityCreateDto)
	return nil
}

// GetExchangeActivityCreateDto ExchangeActivityCreateDto Getter
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetExchangeActivityCreateDto() *ExchangeActivityCreateDto {
	return r._exchangeActivityCreateDto
}

var poolTaobaoCrmExchangeActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmExchangeActivityCreateRequest()
	},
}

// GetTaobaoCrmExchangeActivityCreateRequest 从 sync.Pool 获取 TaobaoCrmExchangeActivityCreateAPIRequest
func GetTaobaoCrmExchangeActivityCreateAPIRequest() *TaobaoCrmExchangeActivityCreateAPIRequest {
	return poolTaobaoCrmExchangeActivityCreateAPIRequest.Get().(*TaobaoCrmExchangeActivityCreateAPIRequest)
}

// ReleaseTaobaoCrmExchangeActivityCreateAPIRequest 将 TaobaoCrmExchangeActivityCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmExchangeActivityCreateAPIRequest(v *TaobaoCrmExchangeActivityCreateAPIRequest) {
	v.Reset()
	poolTaobaoCrmExchangeActivityCreateAPIRequest.Put(v)
}
