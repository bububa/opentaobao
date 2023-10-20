package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantbrandqueryAPIResponse 品牌查询接口 API返回值
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
type AlibabawdkmerchantbrandqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkmerchantbrandqueryAPIResponseModel
}

// AlibabawdkmerchantbrandqueryAPIResponseModel is 品牌查询接口 成功返回结果
type AlibabawdkmerchantbrandqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_brand_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkmerchantbrandqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
