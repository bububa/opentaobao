package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
寄卖V2订单履约 APIRequest
alibaba.idle.consignmentii.order.perform

寄卖V2订单履约，服务商同步订单信息，驱动交易流转
*/
type AlibabaIdleConsignmentiiOrderPerformRequest struct {
    model.Params

    // 同步参数
    consignmentV2OrderSynDto   *ConsignmentV2OrderSynDto 

}

func NewAlibabaIdleConsignmentiiOrderPerformRequest() *AlibabaIdleConsignmentiiOrderPerformRequest{
    return &AlibabaIdleConsignmentiiOrderPerformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.consignmentii.order.perform"
}

func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleConsignmentiiOrderPerformRequest) SetConsignmentV2OrderSynDto(consignmentV2OrderSynDto *ConsignmentV2OrderSynDto) error {
    r.consignmentV2OrderSynDto = consignmentV2OrderSynDto
    r.Set("consignment_v2_order_syn_dto", consignmentV2OrderSynDto)
    return nil
}

func (r AlibabaIdleConsignmentiiOrderPerformRequest) GetConsignmentV2OrderSynDto() *ConsignmentV2OrderSynDto {
    return r.consignmentV2OrderSynDto
}

