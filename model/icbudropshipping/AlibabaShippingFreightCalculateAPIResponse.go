package icbudropshipping

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabashippingfreightcalculateAPIResponse 阿里巴巴商品运费计算查询接口 API返回值
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
type AlibabashippingfreightcalculateAPIResponse struct {
	model.CommonResponse
	AlibabashippingfreightcalculateAPIResponseModel
}

// AlibabashippingfreightcalculateAPIResponseModel is 阿里巴巴商品运费计算查询接口 成功返回结果
type AlibabashippingfreightcalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shipping_freight_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pojo
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
}
