package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleShippinglineTemplateInitAPIRequest 创建初始模板 API请求
// alibaba.wholesale.shippingline.template.init
//
// 创建默认的几种运费模板
type AlibabaWholesaleShippinglineTemplateInitAPIRequest struct {
	model.Params
	// 创建初始运费模板参数
	_initialTemplate []InitialTemplate
}

// NewAlibabaWholesaleShippinglineTemplateInitRequest 初始化AlibabaWholesaleShippinglineTemplateInitAPIRequest对象
func NewAlibabaWholesaleShippinglineTemplateInitRequest() *AlibabaWholesaleShippinglineTemplateInitAPIRequest {
	return &AlibabaWholesaleShippinglineTemplateInitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleShippinglineTemplateInitAPIRequest) GetApiMethodName() string {
	return "alibaba.wholesale.shippingline.template.init"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleShippinglineTemplateInitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInitialTemplate is InitialTemplate Setter
// 创建初始运费模板参数
func (r *AlibabaWholesaleShippinglineTemplateInitAPIRequest) SetInitialTemplate(_initialTemplate []InitialTemplate) error {
	r._initialTemplate = _initialTemplate
	r.Set("initial_template", _initialTemplate)
	return nil
}

// GetInitialTemplate InitialTemplate Getter
func (r AlibabaWholesaleShippinglineTemplateInitAPIRequest) GetInitialTemplate() []InitialTemplate {
	return r._initialTemplate
}
