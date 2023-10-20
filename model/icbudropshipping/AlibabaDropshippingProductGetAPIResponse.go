package icbudropshipping

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadropshippingproductgetAPIResponse 阿里巴巴dropshipping 产品信息获取 API返回值
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
type AlibabadropshippingproductgetAPIResponse struct {
	model.CommonResponse
	AlibabadropshippingproductgetAPIResponseModel
}

// AlibabadropshippingproductgetAPIResponseModel is 阿里巴巴dropshipping 产品信息获取 成功返回结果
type AlibabadropshippingproductgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// product pojo
	Value []DistributionSaleProduct `json:"value,omitempty" xml:"value>distribution_sale_product,omitempty"`
}
