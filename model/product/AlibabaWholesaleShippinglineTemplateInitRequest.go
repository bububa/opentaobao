package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建初始模板 API请求
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板
*/
type AlibabaWholesaleShippinglineTemplateInitRequest struct {
    model.Params
    // 创建初始运费模板参数
    _initialTemplate   []InitialTemplate
}

// 初始化AlibabaWholesaleShippinglineTemplateInitRequest对象
func NewAlibabaWholesaleShippinglineTemplateInitRequest() *AlibabaWholesaleShippinglineTemplateInitRequest{
    return &AlibabaWholesaleShippinglineTemplateInitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetApiMethodName() string {
    return "alibaba.wholesale.shippingline.template.init"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InitialTemplate Setter
// 创建初始运费模板参数
func (r *AlibabaWholesaleShippinglineTemplateInitRequest) SetInitialTemplate(_initialTemplate []InitialTemplate) error {
    r._initialTemplate = _initialTemplate
    r.Set("initial_template", _initialTemplate)
    return nil
}

// InitialTemplate Getter
func (r AlibabaWholesaleShippinglineTemplateInitRequest) GetInitialTemplate() []InitialTemplate {
    return r._initialTemplate
}
