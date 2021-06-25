package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建初始模板 APIRequest
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板
*/
type AlibabaWholesaleShippinglineTemplateInitRequest struct {
    model.Params

    // 创建初始运费模板参数
    initialTemplate   []InitialTemplate 

}

func NewAlibabaWholesaleShippinglineTemplateInitRequest() *AlibabaWholesaleShippinglineTemplateInitRequest{
    return &AlibabaWholesaleShippinglineTemplateInitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetApiMethodName() string {
    return "alibaba.wholesale.shippingline.template.init"
}

func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWholesaleShippinglineTemplateInitRequest) SetInitialTemplate(initialTemplate []InitialTemplate) error {
    r.initialTemplate = initialTemplate
    r.Set("initial_template", initialTemplate)
    return nil
}

func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetInitialTemplate() []InitialTemplate {
    return r.initialTemplate
}

