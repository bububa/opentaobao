package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商订单价格修改接口 APIRequest
alibaba.idle.isv.order.adjustprice

闲鱼用户通过授权的服务商修改订单价格和邮费
*/
type AlibabaIdleIsvOrderAdjustpriceRequest struct {
    model.Params

    // 输入参数
    paramAdjustOrderPrice   *IsvAdjustOrderPriceDto 

}

func NewAlibabaIdleIsvOrderAdjustpriceRequest() *AlibabaIdleIsvOrderAdjustpriceRequest{
    return &AlibabaIdleIsvOrderAdjustpriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.adjustprice"
}

func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvOrderAdjustpriceRequest) SetParamAdjustOrderPrice(paramAdjustOrderPrice *IsvAdjustOrderPriceDto) error {
    r.paramAdjustOrderPrice = paramAdjustOrderPrice
    r.Set("param_adjust_order_price", paramAdjustOrderPrice)
    return nil
}

func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetParamAdjustOrderPrice() *IsvAdjustOrderPriceDto {
    return r.paramAdjustOrderPrice
}

