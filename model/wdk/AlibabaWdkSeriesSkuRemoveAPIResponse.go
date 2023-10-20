package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSkuRemoveAPIResponse 系列品商品变更-移除商品 API返回值
// alibaba.wdk.series.sku.remove
//
// 系列品商品变更-移除商品
type AlibabaWdkSeriesSkuRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSeriesSkuRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSkuRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSeriesSkuRemoveAPIResponseModel).Reset()
}

// AlibabaWdkSeriesSkuRemoveAPIResponseModel is 系列品商品变更-移除商品 成功返回结果
type AlibabaWdkSeriesSkuRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_sku_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabaWdkSeriesSkuRemoveApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSeriesSkuRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkSeriesSkuRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesSkuRemoveAPIResponse)
	},
}

// GetAlibabaWdkSeriesSkuRemoveAPIResponse 从 sync.Pool 获取 AlibabaWdkSeriesSkuRemoveAPIResponse
func GetAlibabaWdkSeriesSkuRemoveAPIResponse() *AlibabaWdkSeriesSkuRemoveAPIResponse {
	return poolAlibabaWdkSeriesSkuRemoveAPIResponse.Get().(*AlibabaWdkSeriesSkuRemoveAPIResponse)
}

// ReleaseAlibabaWdkSeriesSkuRemoveAPIResponse 将 AlibabaWdkSeriesSkuRemoveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSeriesSkuRemoveAPIResponse(v *AlibabaWdkSeriesSkuRemoveAPIResponse) {
	v.Reset()
	poolAlibabaWdkSeriesSkuRemoveAPIResponse.Put(v)
}
