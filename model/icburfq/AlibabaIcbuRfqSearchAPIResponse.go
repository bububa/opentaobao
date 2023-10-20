package icburfq

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqSearchAPIResponse 查询RFQ API返回值
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
type AlibabaIcbuRfqSearchAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuRfqSearchAPIResponseModel).Reset()
}

// AlibabaIcbuRfqSearchAPIResponseModel is 查询RFQ 成功返回结果
type AlibabaIcbuRfqSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfq_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息结果集
	Result *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuRfqSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuRfqSearchAPIResponse)
	},
}

// GetAlibabaIcbuRfqSearchAPIResponse 从 sync.Pool 获取 AlibabaIcbuRfqSearchAPIResponse
func GetAlibabaIcbuRfqSearchAPIResponse() *AlibabaIcbuRfqSearchAPIResponse {
	return poolAlibabaIcbuRfqSearchAPIResponse.Get().(*AlibabaIcbuRfqSearchAPIResponse)
}

// ReleaseAlibabaIcbuRfqSearchAPIResponse 将 AlibabaIcbuRfqSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuRfqSearchAPIResponse(v *AlibabaIcbuRfqSearchAPIResponse) {
	v.Reset()
	poolAlibabaIcbuRfqSearchAPIResponse.Put(v)
}
