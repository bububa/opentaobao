package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesDefaultskuResetAPIResponse 系列品商品变更-重置默认商品 API返回值
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
type AlibabaWdkSeriesDefaultskuResetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesDefaultskuResetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesDefaultskuResetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesDefaultskuResetAPIResponseModel).Reset()
}

// AlibabaWdkSeriesDefaultskuResetAPIResponseModel is 系列品商品变更-重置默认商品 成功返回结果
type AlibabaWdkSeriesDefaultskuResetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_defaultsku_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesDefaultskuResetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesDefaultskuResetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesDefaultskuResetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesDefaultskuResetAPIResponse)
	},
}

// GetAlibabaWdkSeriesDefaultskuResetAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesDefaultskuResetAPIResponse
func GetAlibabaWdkSeriesDefaultskuResetAPIResponse() *AlibabaWdkSeriesDefaultskuResetAPIResponse {
	return poolAlibabaWdkSeriesDefaultskuResetAPIResponse.Get().(*AlibabaWdkSeriesDefaultskuResetAPIResponse)
}

// ReleaseAlibabaWdkSeriesDefaultskuResetAPIResponse 将 AlibabaWdkSeriesDefaultskuResetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesDefaultskuResetAPIResponse(v *AlibabaWdkSeriesDefaultskuResetAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesDefaultskuResetAPIResponse.Put(v)
}
