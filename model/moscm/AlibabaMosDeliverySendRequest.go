package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货 APIRequest
alibaba.mos.delivery.send

订单发货填写快递单
*/
type AlibabaMosDeliverySendRequest struct {
    model.Params

    // 发货信息
    deliveryDto   *DeliveryDto 

}

func NewAlibabaMosDeliverySendRequest() *AlibabaMosDeliverySendRequest{
    return &AlibabaMosDeliverySendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosDeliverySendRequest) GetApiMethodName() string {
    return "alibaba.mos.delivery.send"
}

func (r AlibabaMosDeliverySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosDeliverySendRequest) SetDeliveryDto(deliveryDto *DeliveryDto) error {
    r.deliveryDto = deliveryDto
    r.Set("delivery_dto", deliveryDto)
    return nil
}

func (r AlibabaMosDeliverySendRequest) GetDeliveryDto() *DeliveryDto {
    return r.deliveryDto
}

