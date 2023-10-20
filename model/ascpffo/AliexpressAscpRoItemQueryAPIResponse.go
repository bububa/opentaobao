package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpRoItemQueryAPIResponse AliExpress退供单明细查询API API返回值
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
type AliexpressAscpRoItemQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpRoItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpRoItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpRoItemQueryAPIResponseModel).Reset()
}

// AliexpressAscpRoItemQueryAPIResponseModel is AliExpress退供单明细查询API 成功返回结果
type AliexpressAscpRoItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ro_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dto
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpRoItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpRoItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpRoItemQueryAPIResponse)
	},
}

// GetAliexpressAscpRoItemQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpRoItemQueryAPIResponse
func GetAliexpressAscpRoItemQueryAPIResponse() *AliexpressAscpRoItemQueryAPIResponse {
	return poolAliexpressAscpRoItemQueryAPIResponse.Get().(*AliexpressAscpRoItemQueryAPIResponse)
}

// ReleaseAliexpressAscpRoItemQueryAPIResponse 将 AliexpressAscpRoItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpRoItemQueryAPIResponse(v *AliexpressAscpRoItemQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpRoItemQueryAPIResponse.Put(v)
}
