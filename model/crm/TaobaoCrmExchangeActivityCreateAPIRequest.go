package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建积分兑换活动 API请求
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动
*/
type TaobaoCrmExchangeActivityCreateAPIRequest struct {
    model.Params
    // 创建积分兑换活动
    _exchangeActivityCreateDto   *ExchangeActivityCreateDto
}

// 初始化TaobaoCrmExchangeActivityCreateAPIRequest对象
func NewTaobaoCrmExchangeActivityCreateRequest() *TaobaoCrmExchangeActivityCreateAPIRequest{
    return &TaobaoCrmExchangeActivityCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExchangeActivityCreateDto Setter
// 创建积分兑换活动
func (r *TaobaoCrmExchangeActivityCreateAPIRequest) SetExchangeActivityCreateDto(_exchangeActivityCreateDto *ExchangeActivityCreateDto) error {
    r._exchangeActivityCreateDto = _exchangeActivityCreateDto
    r.Set("exchange_activity_create_dto", _exchangeActivityCreateDto)
    return nil
}

// ExchangeActivityCreateDto Getter
func (r TaobaoCrmExchangeActivityCreateAPIRequest) GetExchangeActivityCreateDto() *ExchangeActivityCreateDto {
    return r._exchangeActivityCreateDto
}
