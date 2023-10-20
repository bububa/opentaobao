package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpOnwayInventoryQueryAPIResponse AliExpress在途库存查询API API返回值
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
type AliexpressAscpOnwayInventoryQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpOnwayInventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpOnwayInventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpOnwayInventoryQueryAPIResponseModel).Reset()
}

// AliexpressAscpOnwayInventoryQueryAPIResponseModel is AliExpress在途库存查询API 成功返回结果
type AliexpressAscpOnwayInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_onway_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressAscpOnwayInventoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpOnwayInventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpOnwayInventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpOnwayInventoryQueryAPIResponse)
	},
}

// GetAliexpressAscpOnwayInventoryQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpOnwayInventoryQueryAPIResponse
func GetAliexpressAscpOnwayInventoryQueryAPIResponse() *AliexpressAscpOnwayInventoryQueryAPIResponse {
	return poolAliexpressAscpOnwayInventoryQueryAPIResponse.Get().(*AliexpressAscpOnwayInventoryQueryAPIResponse)
}

// ReleaseAliexpressAscpOnwayInventoryQueryAPIResponse 将 AliexpressAscpOnwayInventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpOnwayInventoryQueryAPIResponse(v *AliexpressAscpOnwayInventoryQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpOnwayInventoryQueryAPIResponse.Put(v)
}
