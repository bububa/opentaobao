package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpPoQueryAPIResponse AliExpress采购单查询API API返回值
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
type AliexpressAscpPoQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpPoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpPoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpPoQueryAPIResponseModel).Reset()
}

// AliexpressAscpPoQueryAPIResponseModel is AliExpress采购单查询API 成功返回结果
type AliexpressAscpPoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_po_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AliexpressAscpPoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpPoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpPoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpPoQueryAPIResponse)
	},
}

// GetAliexpressAscpPoQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpPoQueryAPIResponse
func GetAliexpressAscpPoQueryAPIResponse() *AliexpressAscpPoQueryAPIResponse {
	return poolAliexpressAscpPoQueryAPIResponse.Get().(*AliexpressAscpPoQueryAPIResponse)
}

// ReleaseAliexpressAscpPoQueryAPIResponse 将 AliexpressAscpPoQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpPoQueryAPIResponse(v *AliexpressAscpPoQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpPoQueryAPIResponse.Put(v)
}
