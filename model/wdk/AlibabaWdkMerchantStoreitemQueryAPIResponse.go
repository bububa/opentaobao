package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantstoreitemqueryAPIResponse 门店商品信心查询 API返回值
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
type AlibabawdkmerchantstoreitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkmerchantstoreitemqueryAPIResponseModel
}

// AlibabawdkmerchantstoreitemqueryAPIResponseModel is 门店商品信心查询 成功返回结果
type AlibabawdkmerchantstoreitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkmerchantstoreitemqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
