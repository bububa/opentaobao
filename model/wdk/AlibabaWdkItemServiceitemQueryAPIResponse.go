package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemServiceitemQueryAPIResponse 查询服务商品 API返回值
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
type AlibabaWdkItemServiceitemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemServiceitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemServiceitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemServiceitemQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemServiceitemQueryAPIResponseModel is 查询服务商品 成功返回结果
type AlibabaWdkItemServiceitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_serviceitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemServiceitemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemServiceitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemServiceitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemServiceitemQueryAPIResponse)
	},
}

// GetAlibabaWdkItemServiceitemQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemServiceitemQueryAPIResponse
func GetAlibabaWdkItemServiceitemQueryAPIResponse() *AlibabaWdkItemServiceitemQueryAPIResponse {
	return poolAlibabaWdkItemServiceitemQueryAPIResponse.Get().(*AlibabaWdkItemServiceitemQueryAPIResponse)
}

// ReleaseAlibabaWdkItemServiceitemQueryAPIResponse 将 AlibabaWdkItemServiceitemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemServiceitemQueryAPIResponse(v *AlibabaWdkItemServiceitemQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemServiceitemQueryAPIResponse.Put(v)
}
