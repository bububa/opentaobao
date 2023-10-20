package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFfoItemQueryAPIResponse AliExpress发货单明细分页查询API API返回值
// aliexpress.ascp.ffo.item.query
//
// AE履约发货单明细分页查询
type AliexpressAscpFfoItemQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpFfoItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpFfoItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpFfoItemQueryAPIResponseModel).Reset()
}

// AliexpressAscpFfoItemQueryAPIResponseModel is AliExpress发货单明细分页查询API 成功返回结果
type AliexpressAscpFfoItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ffo_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressAscpFfoItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpFfoItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpFfoItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpFfoItemQueryAPIResponse)
	},
}

// GetAliexpressAscpFfoItemQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpFfoItemQueryAPIResponse
func GetAliexpressAscpFfoItemQueryAPIResponse() *AliexpressAscpFfoItemQueryAPIResponse {
	return poolAliexpressAscpFfoItemQueryAPIResponse.Get().(*AliexpressAscpFfoItemQueryAPIResponse)
}

// ReleaseAliexpressAscpFfoItemQueryAPIResponse 将 AliexpressAscpFfoItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpFfoItemQueryAPIResponse(v *AliexpressAscpFfoItemQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpFfoItemQueryAPIResponse.Put(v)
}
