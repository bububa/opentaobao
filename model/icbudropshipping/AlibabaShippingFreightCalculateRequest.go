package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴商品运费计算查询接口 APIRequest
alibaba.shipping.freight.calculate

阿里巴巴商品运费计算查询接口
*/
type AlibabaShippingFreightCalculateRequest struct {
    model.Params

    // {}
    paramFreightTemplateRequest   *FreightTemplateRequest 

}

func NewAlibabaShippingFreightCalculateRequest() *AlibabaShippingFreightCalculateRequest{
    return &AlibabaShippingFreightCalculateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaShippingFreightCalculateRequest) GetApiMethodName() string {
    return "alibaba.shipping.freight.calculate"
}

func (r AlibabaShippingFreightCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaShippingFreightCalculateRequest) SetParamFreightTemplateRequest(paramFreightTemplateRequest *FreightTemplateRequest) error {
    r.paramFreightTemplateRequest = paramFreightTemplateRequest
    r.Set("param_freight_template_request", paramFreightTemplateRequest)
    return nil
}

func (r AlibabaShippingFreightCalculateRequest) GetParamFreightTemplateRequest() *FreightTemplateRequest {
    return r.paramFreightTemplateRequest
}

