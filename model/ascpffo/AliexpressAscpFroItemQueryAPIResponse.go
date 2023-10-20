package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFroItemQueryAPIResponse AliExpress销退单明细查询API API返回值
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
type AliexpressAscpFroItemQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpFroItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpFroItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpFroItemQueryAPIResponseModel).Reset()
}

// AliexpressAscpFroItemQueryAPIResponseModel is AliExpress销退单明细查询API 成功返回结果
type AliexpressAscpFroItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_fro_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressAscpFroItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpFroItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpFroItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpFroItemQueryAPIResponse)
	},
}

// GetAliexpressAscpFroItemQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpFroItemQueryAPIResponse
func GetAliexpressAscpFroItemQueryAPIResponse() *AliexpressAscpFroItemQueryAPIResponse {
	return poolAliexpressAscpFroItemQueryAPIResponse.Get().(*AliexpressAscpFroItemQueryAPIResponse)
}

// ReleaseAliexpressAscpFroItemQueryAPIResponse 将 AliexpressAscpFroItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpFroItemQueryAPIResponse(v *AliexpressAscpFroItemQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpFroItemQueryAPIResponse.Put(v)
}
