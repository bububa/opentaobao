package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuSearchAPIResponse 供应链中台货品搜索接口 API返回值
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
type AlibabaAscpCnskuSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAscpCnskuSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpCnskuSearchAPIResponseModel).Reset()
}

// AlibabaAscpCnskuSearchAPIResponseModel is 供应链中台货品搜索接口 成功返回结果
type AlibabaAscpCnskuSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_cnsku_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CnskuResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpCnskuSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpCnskuSearchAPIResponse)
	},
}

// GetAlibabaAscpCnskuSearchAPIResponse 从 sync.Pool 获取 AlibabaAscpCnskuSearchAPIResponse
func GetAlibabaAscpCnskuSearchAPIResponse() *AlibabaAscpCnskuSearchAPIResponse {
	return poolAlibabaAscpCnskuSearchAPIResponse.Get().(*AlibabaAscpCnskuSearchAPIResponse)
}

// ReleaseAlibabaAscpCnskuSearchAPIResponse 将 AlibabaAscpCnskuSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpCnskuSearchAPIResponse(v *AlibabaAscpCnskuSearchAPIResponse) {
	v.Reset()
	poolAlibabaAscpCnskuSearchAPIResponse.Put(v)
}
