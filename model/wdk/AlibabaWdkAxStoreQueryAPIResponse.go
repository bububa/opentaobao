package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreQueryAPIResponse 翱象经营店查询接口 API返回值
// alibaba.wdk.ax.store.query
//
// 翱象经营店查询接口
type AlibabaWdkAxStoreQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkAxStoreQueryAPIResponseModel
}

// AlibabaWdkAxStoreQueryAPIResponseModel is 翱象经营店查询接口 成功返回结果
type AlibabaWdkAxStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ax_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询接口返回结果
	ApiResult *AlibabaWdkAxStoreQueryApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
