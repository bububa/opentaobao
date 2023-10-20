package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpPoItemQueryAPIResponse AliExpress采购单明细查询API API返回值
// aliexpress.ascp.po.item.query
//
// AE 供应链仓发 采购单明细查询
type AliexpressAscpPoItemQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpPoItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpPoItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpPoItemQueryAPIResponseModel).Reset()
}

// AliexpressAscpPoItemQueryAPIResponseModel is AliExpress采购单明细查询API 成功返回结果
type AliexpressAscpPoItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_po_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// demo
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpPoItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpPoItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpPoItemQueryAPIResponse)
	},
}

// GetAliexpressAscpPoItemQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpPoItemQueryAPIResponse
func GetAliexpressAscpPoItemQueryAPIResponse() *AliexpressAscpPoItemQueryAPIResponse {
	return poolAliexpressAscpPoItemQueryAPIResponse.Get().(*AliexpressAscpPoItemQueryAPIResponse)
}

// ReleaseAliexpressAscpPoItemQueryAPIResponse 将 AliexpressAscpPoItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpPoItemQueryAPIResponse(v *AliexpressAscpPoItemQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpPoItemQueryAPIResponse.Put(v)
}
