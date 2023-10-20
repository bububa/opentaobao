package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesEditAPIResponse 系列品变更-更新系列 API返回值
// alibaba.wdk.series.edit
//
// 系列品变更-更新系列
type AlibabaWdkSeriesEditAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesEditAPIResponseModel).Reset()
}

// AlibabaWdkSeriesEditAPIResponseModel is 系列品变更-更新系列 成功返回结果
type AlibabaWdkSeriesEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesEditApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesEditAPIResponse)
	},
}

// GetAlibabaWdkSeriesEditAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesEditAPIResponse
func GetAlibabaWdkSeriesEditAPIResponse() *AlibabaWdkSeriesEditAPIResponse {
	return poolAlibabaWdkSeriesEditAPIResponse.Get().(*AlibabaWdkSeriesEditAPIResponse)
}

// ReleaseAlibabaWdkSeriesEditAPIResponse 将 AlibabaWdkSeriesEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesEditAPIResponse(v *AlibabaWdkSeriesEditAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesEditAPIResponse.Put(v)
}
