package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemcategoryqueryAPIResponse 类目查询接口 API返回值
// alibaba.wdk.item.category.query
//
// 类目查询接口
type AlibabawdkitemcategoryqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemcategoryqueryAPIResponseModel
}

// AlibabawdkitemcategoryqueryAPIResponseModel is 类目查询接口 成功返回结果
type AlibabawdkitemcategoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkitemcategoryqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
