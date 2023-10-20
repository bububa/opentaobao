package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskuQueryAPIResponse 门店商品信息查询 API返回值
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
type AlibabaWdkItemStoreskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemStoreskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemStoreskuQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemStoreskuQueryAPIResponseModel is 门店商品信息查询 成功返回结果
type AlibabaWdkItemStoreskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storesku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemStoreskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemStoreskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskuQueryAPIResponse)
	},
}

// GetAlibabaWdkItemStoreskuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemStoreskuQueryAPIResponse
func GetAlibabaWdkItemStoreskuQueryAPIResponse() *AlibabaWdkItemStoreskuQueryAPIResponse {
	return poolAlibabaWdkItemStoreskuQueryAPIResponse.Get().(*AlibabaWdkItemStoreskuQueryAPIResponse)
}

// ReleaseAlibabaWdkItemStoreskuQueryAPIResponse 将 AlibabaWdkItemStoreskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemStoreskuQueryAPIResponse(v *AlibabaWdkItemStoreskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemStoreskuQueryAPIResponse.Put(v)
}
