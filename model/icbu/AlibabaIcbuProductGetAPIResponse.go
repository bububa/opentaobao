package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductgetAPIResponse 获得单个商品详情 API返回值
// alibaba.icbu.product.get
//
// 获取商品详情
type AlibabaicbuproductgetAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproductgetAPIResponseModel
}

// AlibabaicbuproductgetAPIResponseModel is 获得单个商品详情 成功返回结果
type AlibabaicbuproductgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单个商品详情
	Product *AlibabaProductResponse `json:"product,omitempty" xml:"product,omitempty"`
}
