package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
寄卖V2订单履约 API请求
alibaba.idle.consignmentii.order.perform

寄卖V2订单履约，服务商同步订单信息，驱动交易流转
*/
type AlibabaIdleConsignmentiiOrderPerformRequest struct {
    model.Params
    // 同步参数
    _consignmentV2OrderSynDto   *ConsignmentV2OrderSynDto
}

// 初始化AlibabaIdleConsignmentiiOrderPerformRequest对象
func NewAlibabaIdleConsignmentiiOrderPerformRequest() *AlibabaIdleConsignmentiiOrderPerformRequest{
    return &AlibabaIdleConsignmentiiOrderPerformRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.consignmentii.order.perform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignmentV2OrderSynDto Setter
// 同步参数
func (r *AlibabaIdleConsignmentiiOrderPerformRequest) SetConsignmentV2OrderSynDto(_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto) error {
    r._consignmentV2OrderSynDto = _consignmentV2OrderSynDto
    r.Set("consignment_v2_order_syn_dto", _consignmentV2OrderSynDto)
    return nil
}

// ConsignmentV2OrderSynDto Getter
func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetConsignmentV2OrderSynDto() *ConsignmentV2OrderSynDto {
    return r._consignmentV2OrderSynDto
}
