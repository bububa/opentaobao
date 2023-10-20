package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuqueryAPIResponse 查询渠道商品 API返回值
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
type AlibabawdkskuchannelskuqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkskuchannelskuqueryAPIResponseModel
}

// AlibabawdkskuchannelskuqueryAPIResponseModel is 查询渠道商品 成功返回结果
type AlibabawdkskuchannelskuqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskuchannelskuqueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
