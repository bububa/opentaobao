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
type AlibabaOrderFreightCalculateAPIRequest struct {
    model.Params
    // {}
    _paramMultiFreightTemplateRequest   *MultiFreightTemplateRequest
}

// 初始化AlibabaOrderFreightCalculateAPIRequest对象
func NewAlibabaOrderFreightCalculateRequest() *AlibabaOrderFreightCalculateAPIRequest{
    return &AlibabaOrderFreightCalculateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOrderFreightCalculateAPIRequest) GetApiMethodName() string {
    return "alibaba.order.freight.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOrderFreightCalculateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamMultiFreightTemplateRequest Setter
// {}
func (r *AlibabaOrderFreightCalculateAPIRequest) SetParamMultiFreightTemplateRequest(_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest) error {
    r._paramMultiFreightTemplateRequest = _paramMultiFreightTemplateRequest
    r.Set("param_multi_freight_template_request", _paramMultiFreightTemplateRequest)
    return nil
}

// ParamMultiFreightTemplateRequest Getter
func (r AlibabaOrderFreightCalculateAPIRequest) GetParamMultiFreightTemplateRequest() *MultiFreightTemplateRequest {
    return r._paramMultiFreightTemplateRequest
}
