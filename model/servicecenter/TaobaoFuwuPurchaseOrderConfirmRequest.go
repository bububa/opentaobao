package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务市场内购服务下单接口 APIRequest
taobao.fuwu.purchase.order.confirm

通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
*/
type TaobaoFuwuPurchaseOrderConfirmRequest struct {
    model.Params

    // 内购服务下单接口参数
    paramOrderConfirmQueryDTO   *OrderConfirmQueryDto 

}

func NewTaobaoFuwuPurchaseOrderConfirmRequest() *TaobaoFuwuPurchaseOrderConfirmRequest{
    return &TaobaoFuwuPurchaseOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFuwuPurchaseOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.fuwu.purchase.order.confirm"
}

func (r TaobaoFuwuPurchaseOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFuwuPurchaseOrderConfirmRequest) SetParamOrderConfirmQueryDTO(paramOrderConfirmQueryDTO *OrderConfirmQueryDto) error {
    r.paramOrderConfirmQueryDTO = paramOrderConfirmQueryDTO
    r.Set("param_order_confirm_query_d_t_o", paramOrderConfirmQueryDTO)
    return nil
}

func (r TaobaoFuwuPurchaseOrderConfirmRequest) GetParamOrderConfirmQueryDTO() *OrderConfirmQueryDto {
    return r.paramOrderConfirmQueryDTO
}

