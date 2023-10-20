package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpRoQueryAPIResponse AliExpress退供单查询API API返回值
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
type AliexpressAscpRoQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpRoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpRoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpRoQueryAPIResponseModel).Reset()
}

// AliexpressAscpRoQueryAPIResponseModel is AliExpress退供单查询API 成功返回结果
type AliexpressAscpRoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_ro_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressAscpRoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpRoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpRoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpRoQueryAPIResponse)
	},
}

// GetAliexpressAscpRoQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpRoQueryAPIResponse
func GetAliexpressAscpRoQueryAPIResponse() *AliexpressAscpRoQueryAPIResponse {
	return poolAliexpressAscpRoQueryAPIResponse.Get().(*AliexpressAscpRoQueryAPIResponse)
}

// ReleaseAliexpressAscpRoQueryAPIResponse 将 AliexpressAscpRoQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpRoQueryAPIResponse(v *AliexpressAscpRoQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpRoQueryAPIResponse.Put(v)
}
