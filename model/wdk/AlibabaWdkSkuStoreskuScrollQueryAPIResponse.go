package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskustoreskuscrollqueryAPIResponse 门店商品批量查询接口 API返回值
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
type AlibabawdkskustoreskuscrollqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdkskustoreskuscrollqueryAPIResponseModel
}

// AlibabawdkskustoreskuscrollqueryAPIResponseModel is 门店商品批量查询接口 成功返回结果
type AlibabawdkskustoreskuscrollqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_storesku_scroll_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
