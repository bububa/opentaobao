package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkshopqueryAPIResponse 门店查询接口 API返回值
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
type AlibabawdkshopqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkshopqueryAPIResponseModel
}

// AlibabawdkshopqueryAPIResponseModel is 门店查询接口 成功返回结果
type AlibabawdkshopqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_shop_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkshopqueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
