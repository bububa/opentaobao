package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSortAPIResponse 系列品-商品排序 API返回值
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
type AlibabaWdkSeriesSortAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesSortAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSortAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesSortAPIResponseModel).Reset()
}

// AlibabaWdkSeriesSortAPIResponseModel is 系列品-商品排序 成功返回结果
type AlibabaWdkSeriesSortAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_sort_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesSortApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSortAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesSortAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesSortAPIResponse)
	},
}

// GetAlibabaWdkSeriesSortAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesSortAPIResponse
func GetAlibabaWdkSeriesSortAPIResponse() *AlibabaWdkSeriesSortAPIResponse {
	return poolAlibabaWdkSeriesSortAPIResponse.Get().(*AlibabaWdkSeriesSortAPIResponse)
}

// ReleaseAlibabaWdkSeriesSortAPIResponse 将 AlibabaWdkSeriesSortAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesSortAPIResponse(v *AlibabaWdkSeriesSortAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesSortAPIResponse.Put(v)
}
