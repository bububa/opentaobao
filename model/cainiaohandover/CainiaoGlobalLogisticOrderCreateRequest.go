package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流订单 APIRequest
cainiao.global.logistic.order.create

创建物流订单
*/
type CainiaoGlobalLogisticOrderCreateRequest struct {
    model.Params

    // 订单参数
    orderParam   *OpenOrderParam 

    // 多语言
    locale   string 

}

func NewCainiaoGlobalLogisticOrderCreateRequest() *CainiaoGlobalLogisticOrderCreateRequest{
    return &CainiaoGlobalLogisticOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalLogisticOrderCreateRequest) GetApiMethodName() string {
    return "cainiao.global.logistic.order.create"
}

func (r CainiaoGlobalLogisticOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalLogisticOrderCreateRequest) SetOrderParam(orderParam *OpenOrderParam) error {
    r.orderParam = orderParam
    r.Set("order_param", orderParam)
    return nil
}

func (r CainiaoGlobalLogisticOrderCreateRequest) GetOrderParam() *OpenOrderParam {
    return r.orderParam
}

func (r *CainiaoGlobalLogisticOrderCreateRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalLogisticOrderCreateRequest) GetLocale() string {
    return r.locale
}

