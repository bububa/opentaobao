package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDropshippingProductGetAPIRequest
阿里巴巴dropshipping 产品信息获取 API请求
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取 */
type AlibabaDropshippingProductGetAPIRequest struct {
	model.Params
	// {}
	_paramDistributionSaleProductRequest *DistributionSaleProductRequest
}

// New
