package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesEditAPIResponse
系列品变更-更新系列 API返回值
alibaba.wdk.series.edit

系列品变更-更新系列 */
type AlibabaWdkSeriesEditAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesEditAPIResponseModel
}

// AlibabaWdkSeriesEditAPIResponseModel is 系列品变更-更新系列 成功返回结果
type AlibabaWdkSeriesEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesEditApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
