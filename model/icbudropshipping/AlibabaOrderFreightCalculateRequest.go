package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴下单场景运费方案计算 API请求
alibaba.order.freight.calculate

icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
alibaba Create order scenario freight calculation
*/
type AlibabaOrderFreightCalculateRequest struct {
    model.Params
    // {}
    paramMultiFreightTemplateRequest   *MultiFreightTemplateRequest
}

// 初始化AlibabaOrderFreightCalculateRequest对象
func NewAlibabaOrderFreightCalculateRequest() *AlibabaOrderFreightCalculateRequest{
    return &AlibabaOrderFreightCalculateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOrderFreightCalculateRequest) GetApiMethodName() string {
    return "alibaba.order.freight.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOrderFreightCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamMultiFreightTemplateRequest Setter
// {}
func (r *AlibabaOrderFreightCalculateRequest) SetParamMultiFreightTemplateRequest(paramMultiFreightTemplateRequest *MultiFreightTemplateRequest) error {
    r.paramMultiFreightTemplateRequest = paramMultiFreightTemplateRequest
    r.Set("param_multi_freight_template_request", paramMultiFreightTemplateRequest)
    return nil
}

// ParamMultiFreightTemplateRequest Getter
func (r AlibabaOrderFreightCalculateRequest) GetParamMultiFreightTemplateRequest() *MultiFreightTemplateRequest {
    return r.paramMultiFreightTemplateRequest
}
