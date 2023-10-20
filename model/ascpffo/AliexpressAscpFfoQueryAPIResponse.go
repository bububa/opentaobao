package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFfoQueryAPIResponse AliExpress发货单查询API API返回值
// aliexpress.ascp.ffo.query
//
// AE 履约发货单分页查询接口
type AliexpressAscpFfoQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpFfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpFfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpFfoQueryAPIResponseModel).Reset()
}

// AliexpressAscpFfoQueryAPIResponseModel is AliExpress发货单查询API 成功返回结果
type AliexpressAscpFfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ffo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dto
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpFfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpFfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpFfoQueryAPIResponse)
	},
}

// GetAliexpressAscpFfoQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpFfoQueryAPIResponse
func GetAliexpressAscpFfoQueryAPIResponse() *AliexpressAscpFfoQueryAPIResponse {
	return poolAliexpressAscpFfoQueryAPIResponse.Get().(*AliexpressAscpFfoQueryAPIResponse)
}

// ReleaseAliexpressAscpFfoQueryAPIResponse 将 AliexpressAscpFfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpFfoQueryAPIResponse(v *AliexpressAscpFfoQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpFfoQueryAPIResponse.Put(v)
}
