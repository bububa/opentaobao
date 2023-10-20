package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantBrandQueryAPIResponse 品牌查询接口 API返回值
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
type AlibabaWdkMerchantBrandQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantBrandQueryAPIResponseModel
}

// AlibabaWdkMerchantBrandQueryAPIResponseModel is 品牌查询接口 成功返回结果
type AlibabaWdkMerchantBrandQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_brand_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMerchantBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
