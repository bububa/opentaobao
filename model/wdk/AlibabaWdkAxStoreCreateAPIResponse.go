package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreCreateAPIResponse 翱象经营店创建接口 API返回值
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
type AlibabaWdkAxStoreCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkAxStoreCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkAxStoreCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkAxStoreCreateAPIResponseModel).Reset()
}

// AlibabaWdkAxStoreCreateAPIResponseModel is 翱象经营店创建接口 成功返回结果
type AlibabaWdkAxStoreCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ax_store_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkAxStoreCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkAxStoreCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkAxStoreCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkAxStoreCreateAPIResponse)
	},
}

// GetAlibabaWdkAxStoreCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkAxStoreCreateAPIResponse
func GetAlibabaWdkAxStoreCreateAPIResponse() *AlibabaWdkAxStoreCreateAPIResponse {
	return poolAlibabaWdkAxStoreCreateAPIResponse.Get().(*AlibabaWdkAxStoreCreateAPIResponse)
}

// ReleaseAlibabaWdkAxStoreCreateAPIResponse 将 AlibabaWdkAxStoreCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkAxStoreCreateAPIResponse(v *AlibabaWdkAxStoreCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkAxStoreCreateAPIResponse.Put(v)
}
