package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供给买家使用的运费计算接口 APIRequest
aliexpress.logistics.buyer.freight.calculate

提供给买家使用的运费计算接口
*/
type AliexpressLogisticsBuyerFreightCalculateRequest struct {
    model.Params

    // 运费计算请求参数
    paramAeopFreightCalculateForBuyerDTO   *AeopFreightCalculateForBuyerDto 

}

func NewAliexpressLogisticsBuyerFreightCalculateRequest() *AliexpressLogisticsBuyerFreightCalculateRequest{
    return &AliexpressLogisticsBuyerFreightCalculateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressLogisticsBuyerFreightCalculateRequest) GetApiMethodName() string {
    return "aliexpress.logistics.buyer.freight.calculate"
}

func (r AliexpressLogisticsBuyerFreightCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressLogisticsBuyerFreightCalculateRequest) SetParamAeopFreightCalculateForBuyerDTO(paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto) error {
    r.paramAeopFreightCalculateForBuyerDTO = paramAeopFreightCalculateForBuyerDTO
    r.Set("param_aeop_freight_calculate_for_buyer_d_t_o", paramAeopFreightCalculateForBuyerDTO)
    return nil
}

func (r AliexpressLogisticsBuyerFreightCalculateRequest) GetParamAeopFreightCalculateForBuyerDTO() *AeopFreightCalculateForBuyerDto {
    return r.paramAeopFreightCalculateForBuyerDTO
}

