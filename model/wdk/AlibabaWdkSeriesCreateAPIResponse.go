package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesCreateAPIResponse 系列品变更-新增系列 API返回值
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
type AlibabaWdkSeriesCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesCreateAPIResponseModel
}

// AlibabaWdkSeriesCreateAPIResponseModel is 系列品变更-新增系列 成功返回结果
type AlibabaWdkSeriesCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
