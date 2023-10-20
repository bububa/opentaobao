package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpInventoryLogQueryAPIResponse AliExpress库存流水查询API API返回值
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
type AliexpressAscpInventoryLogQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpInventoryLogQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpInventoryLogQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpInventoryLogQueryAPIResponseModel).Reset()
}

// AliexpressAscpInventoryLogQueryAPIResponseModel is AliExpress库存流水查询API 成功返回结果
type AliexpressAscpInventoryLogQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_inventory_log_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询出参
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpInventoryLogQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpInventoryLogQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpInventoryLogQueryAPIResponse)
	},
}

// GetAliexpressAscpInventoryLogQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpInventoryLogQueryAPIResponse
func GetAliexpressAscpInventoryLogQueryAPIResponse() *AliexpressAscpInventoryLogQueryAPIResponse {
	return poolAliexpressAscpInventoryLogQueryAPIResponse.Get().(*AliexpressAscpInventoryLogQueryAPIResponse)
}

// ReleaseAliexpressAscpInventoryLogQueryAPIResponse 将 AliexpressAscpInventoryLogQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpInventoryLogQueryAPIResponse(v *AliexpressAscpInventoryLogQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpInventoryLogQueryAPIResponse.Put(v)
}
