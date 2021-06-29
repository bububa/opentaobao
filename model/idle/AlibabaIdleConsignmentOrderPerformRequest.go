package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
帮卖订单履约 APIRequest
alibaba.idle.consignment.order.perform

帮卖订单履约，回收商同步订单信息，驱动交易流转
*/
type AlibabaIdleConsignmentOrderPerformRequest struct {
    model.Params

    // 帮卖订单同步DTO
    param   *ConsignmentOrderSynDto 

}

func NewAlibabaIdleConsignmentOrderPerformRequest() *AlibabaIdleConsignmentOrderPerformRequest{
    return &AlibabaIdleConsignmentOrderPerformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleConsignmentOrderPerformRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.order.perform"
}

func (r AlibabaIdleConsignmentOrderPerformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleConsignmentOrderPerformRequest) SetParam(param *ConsignmentOrderSynDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleConsignmentOrderPerformRequest) GetParam() *ConsignmentOrderSynDto {
    return r.param
}

