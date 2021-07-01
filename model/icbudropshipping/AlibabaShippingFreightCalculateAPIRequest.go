package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaShippingFreightCalculateAPIRequest
阿里巴巴商品运费计算查询接口 API请求
alibaba.shipping.freight.calculate

阿里巴巴商品运费计算查询接口 */
type AlibabaShippingFreightCalculateAPIRequest struct {
	model.Params
	// {}
	_paramFreightTemplateRequest *FreightTemplateRequest
}

// New
