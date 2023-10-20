package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantitemqueryAPIResponse 商家商品查询 API返回值
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
type AlibabawdkmerchantitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkmerchantitemqueryAPIResponseModel
}

// AlibabawdkmerchantitemqueryAPIResponseModel is 商家商品查询 成功返回结果
type AlibabawdkmerchantitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkmerchantitemqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
