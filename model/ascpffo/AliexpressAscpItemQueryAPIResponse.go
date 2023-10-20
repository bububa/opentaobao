package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpItemQueryAPIResponse AliExpress货品查询查询API API返回值
// aliexpress.ascp.item.query
//
// AE货品查询API
type AliexpressAscpItemQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpItemQueryAPIResponseModel).Reset()
}

// AliexpressAscpItemQueryAPIResponseModel is AliExpress货品查询查询API 成功返回结果
type AliexpressAscpItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// DTO
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpItemQueryAPIResponse)
	},
}

// GetAliexpressAscpItemQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpItemQueryAPIResponse
func GetAliexpressAscpItemQueryAPIResponse() *AliexpressAscpItemQueryAPIResponse {
	return poolAliexpressAscpItemQueryAPIResponse.Get().(*AliexpressAscpItemQueryAPIResponse)
}

// ReleaseAliexpressAscpItemQueryAPIResponse 将 AliexpressAscpItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpItemQueryAPIResponse(v *AliexpressAscpItemQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpItemQueryAPIResponse.Put(v)
}
