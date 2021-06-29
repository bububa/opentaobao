package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货 API请求
alibaba.mos.delivery.send

订单发货填写快递单
*/
type AlibabaMosDeliverySendRequest struct {
    model.Params
    // 发货信息
    _deliveryDto   *DeliveryDTO
}

// 初始化AlibabaMosDeliverySendRequest对象
func NewAlibabaMosDeliverySendRequest() *AlibabaMosDeliverySendRequest{
    return &AlibabaMosDeliverySendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosDeliverySendRequest) GetApiMethodName() string {
    return "alibaba.mos.delivery.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosDeliverySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryDto Setter
// 发货信息
func (r *AlibabaMosDeliverySendRequest) SetDeliveryDto(_deliveryDto *DeliveryDTO) error {
    r._deliveryDto = _deliveryDto
    r.Set("delivery_dto", _deliveryDto)
    return nil
}

// DeliveryDto Getter
func (r AlibabaMosDeliverySendRequest) GetDeliveryDto() *DeliveryDTO {
    return r._deliveryDto
}
