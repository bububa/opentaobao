package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWholesaleShippinglineTemplateInitAPIRequest
创建初始模板 API请求
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板 */
type AlibabaWholesaleShippinglineTemplateInitAPIRequest struct {
	model.Params
	// 创建初始运费模板参数
	_initialTemplate []InitialTemplate
}

// New
