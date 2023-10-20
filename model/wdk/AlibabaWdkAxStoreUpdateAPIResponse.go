package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreUpdateAPIResponse 翱翔经营店更新接口 API返回值
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
type AlibabaWdkAxStoreUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkAxStoreUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkAxStoreUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkAxStoreUpdateAPIResponseModel).Reset()
}

// AlibabaWdkAxStoreUpdateAPIResponseModel is 翱翔经营店更新接口 成功返回结果
type AlibabaWdkAxStoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ax_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用接口返回结果
	ApiResult *AlibabaWdkAxStoreUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkAxStoreUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkAxStoreUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkAxStoreUpdateAPIResponse)
	},
}

// GetAlibabaWdkAxStoreUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkAxStoreUpdateAPIResponse
func GetAlibabaWdkAxStoreUpdateAPIResponse() *AlibabaWdkAxStoreUpdateAPIResponse {
	return poolAlibabaWdkAxStoreUpdateAPIResponse.Get().(*AlibabaWdkAxStoreUpdateAPIResponse)
}

// ReleaseAlibabaWdkAxStoreUpdateAPIResponse 将 AlibabaWdkAxStoreUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkAxStoreUpdateAPIResponse(v *AlibabaWdkAxStoreUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkAxStoreUpdateAPIResponse.Put(v)
}
