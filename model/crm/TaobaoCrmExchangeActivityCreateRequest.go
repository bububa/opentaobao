package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建积分兑换活动 APIRequest
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动
*/
type TaobaoCrmExchangeActivityCreateRequest struct {
    model.Params

    // 创建积分兑换活动
    exchangeActivityCreateDto   *ExchangeActivityCreateDto 

}

func NewTaobaoCrmExchangeActivityCreateRequest() *TaobaoCrmExchangeActivityCreateRequest{
    return &TaobaoCrmExchangeActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmExchangeActivityCreateRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.activity.create"
}

func (r TaobaoCrmExchangeActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmExchangeActivityCreateRequest) SetExchangeActivityCreateDto(exchangeActivityCreateDto *ExchangeActivityCreateDto) error {
    r.exchangeActivityCreateDto = exchangeActivityCreateDto
    r.Set("exchange_activity_create_dto", exchangeActivityCreateDto)
    return nil
}

func (r TaobaoCrmExchangeActivityCreateRequest) GetExchangeActivityCreateDto() *ExchangeActivityCreateDto {
    return r.exchangeActivityCreateDto
}

