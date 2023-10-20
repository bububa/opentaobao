package icburfq

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqRecommendAPIResponse rfq推荐 API返回值
// alibaba.icbu.rfq.recommend
//
// rfq推荐
type AlibabaIcbuRfqRecommendAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqRecommendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqRecommendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuRfqRecommendAPIResponseModel).Reset()
}

// AlibabaIcbuRfqRecommendAPIResponseModel is rfq推荐 成功返回结果
type AlibabaIcbuRfqRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfq_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqRecommendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuRfqRecommendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuRfqRecommendAPIResponse)
	},
}

// GetAlibabaIcbuRfqRecommendAPIResponse 从 sync.Pool 获取 AlibabaIcbuRfqRecommendAPIResponse
func GetAlibabaIcbuRfqRecommendAPIResponse() *AlibabaIcbuRfqRecommendAPIResponse {
	return poolAlibabaIcbuRfqRecommendAPIResponse.Get().(*AlibabaIcbuRfqRecommendAPIResponse)
}

// ReleaseAlibabaIcbuRfqRecommendAPIResponse 将 AlibabaIcbuRfqRecommendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuRfqRecommendAPIResponse(v *AlibabaIcbuRfqRecommendAPIResponse) {
	v.Reset()
	poolAlibabaIcbuRfqRecommendAPIResponse.Put(v)
}
