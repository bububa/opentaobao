package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenVersionCountAPIResponse 版本数量查询 API返回值
// alibaba.wdk.marketing.open.version.count
//
// 版本数量查询
type AlibabaWdkMarketingOpenVersionCountAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenVersionCountAPIResponseModel
}

// AlibabaWdkMarketingOpenVersionCountAPIResponseModel is 版本数量查询 成功返回结果
type AlibabaWdkMarketingOpenVersionCountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_version_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
