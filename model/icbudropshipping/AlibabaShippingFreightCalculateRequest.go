package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴商品运费计算查询接口 API请求
alibaba.shipping.freight.calculate

阿里巴巴商品运费计算查询接口
*/
type AlibabaShippingFreightCalculateRequest struct {
    model.Params
    // {}
    paramFreightTemplateRequest   *FreightTemplateRequest
}

// 初始化AlibabaShippingFreightCalculateRequest对象
func NewAlibabaShippingFreightCalculateRequest() *AlibabaShippingFreightCalculateRequest{
    return &AlibabaShippingFreightCalculateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaShippingFreightCalculateRequest) GetApiMethodName() string {
    return "alibaba.shipping.freight.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaShippingFreightCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamFreightTemplateRequest Setter
// {}
func (r *AlibabaShippingFreightCalculateRequest) SetParamFreightTemplateRequest(paramFreightTemplateRequest *FreightTemplateRequest) error {
    r.paramFreightTemplateRequest = paramFreightTemplateRequest
    r.Set("param_freight_template_request", paramFreightTemplateRequest)
    return nil
}

// ParamFreightTemplateRequest Getter
func (r AlibabaShippingFreightCalculateRequest) GetParamFreightTemplateRequest() *FreightTemplateRequest {
    return r.paramFreightTemplateRequest
}
