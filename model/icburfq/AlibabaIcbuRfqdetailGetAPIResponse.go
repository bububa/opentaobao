package icburfq

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqdetailGetAPIResponse 获取RFQ详情 API返回值
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
type AlibabaIcbuRfqdetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqdetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqdetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuRfqdetailGetAPIResponseModel).Reset()
}

// AlibabaIcbuRfqdetailGetAPIResponseModel is 获取RFQ详情 成功返回结果
type AlibabaIcbuRfqdetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfqdetail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqdetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuRfqdetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuRfqdetailGetAPIResponse)
	},
}

// GetAlibabaIcbuRfqdetailGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuRfqdetailGetAPIResponse
func GetAlibabaIcbuRfqdetailGetAPIResponse() *AlibabaIcbuRfqdetailGetAPIResponse {
	return poolAlibabaIcbuRfqdetailGetAPIResponse.Get().(*AlibabaIcbuRfqdetailGetAPIResponse)
}

// ReleaseAlibabaIcbuRfqdetailGetAPIResponse 将 AlibabaIcbuRfqdetailGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuRfqdetailGetAPIResponse(v *AlibabaIcbuRfqdetailGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuRfqdetailGetAPIResponse.Put(v)
}
