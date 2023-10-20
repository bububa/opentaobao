package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmexchangeactivitycreateAPIRequest 创建积分兑换活动 API请求
// taobao.crm.exchange.activity.create
//
// 创建针对积分兑换类型的活动
type TaobaocrmexchangeactivitycreateAPIRequest struct {
	model.Params
	// 创建积分兑换活动
	_exchangeActivityCreateDto *ExchangeActivityCreateDto
}

// NewTaobaocrmexchangeactivitycreateRequest 初始化TaobaocrmexchangeactivitycreateAPIRequest对象
func NewTaobaocrmexchangeactivitycreateRequest() *TaobaocrmexchangeactivitycreateAPIRequest {
	return &TaobaocrmexchangeactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmexchangeactivitycreateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmexchangeactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmexchangeactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExchangeActivityCreateDto is ExchangeActivityCreateDto Setter
// 创建积分兑换活动
func (r *TaobaocrmexchangeactivitycreateAPIRequest) SetExchangeActivityCreateDto(_exchangeActivityCreateDto *ExchangeActivityCreateDto) error {
	r._exchangeActivityCreateDto = _exchangeActivityCreateDto
	r.Set("exchange_activity_create_dto", _exchangeActivityCreateDto)
	return nil
}

// GetExchangeActivityCreateDto ExchangeActivityCreateDto Getter
func (r TaobaocrmexchangeactivitycreateAPIRequest) GetExchangeActivityCreateDto() *ExchangeActivityCreateDto {
	return r._exchangeActivityCreateDto
}
