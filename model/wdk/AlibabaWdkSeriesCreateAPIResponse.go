package wdk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaWdkSeriesCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesCreateAPIResponseModel).Reset()
}

// AlibabaWdkSeriesCreateAPIResponseModel is 系列品变更-新增系列 成功返回结果
type AlibabaWdkSeriesCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesCreateAPIResponse)
	},
}

// GetAlibabaWdkSeriesCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesCreateAPIResponse
func GetAlibabaWdkSeriesCreateAPIResponse() *AlibabaWdkSeriesCreateAPIResponse {
	return poolAlibabaWdkSeriesCreateAPIResponse.Get().(*AlibabaWdkSeriesCreateAPIResponse)
}

// ReleaseAlibabaWdkSeriesCreateAPIResponse 将 AlibabaWdkSeriesCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesCreateAPIResponse(v *AlibabaWdkSeriesCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesCreateAPIResponse.Put(v)
}
