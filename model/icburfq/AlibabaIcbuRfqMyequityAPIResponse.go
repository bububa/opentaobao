package icburfq

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqMyequityAPIResponse 我的权益 API返回值
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
type AlibabaIcbuRfqMyequityAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqMyequityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqMyequityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuRfqMyequityAPIResponseModel).Reset()
}

// AlibabaIcbuRfqMyequityAPIResponseModel is 我的权益 成功返回结果
type AlibabaIcbuRfqMyequityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfq_myequity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuRfqMyequityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaIcbuRfqMyequityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuRfqMyequityAPIResponse)
	},
}

// GetAlibabaIcbuRfqMyequityAPIResponse 从 sync.Pool 获取 AlibabaIcbuRfqMyequityAPIResponse
func GetAlibabaIcbuRfqMyequityAPIResponse() *AlibabaIcbuRfqMyequityAPIResponse {
	return poolAlibabaIcbuRfqMyequityAPIResponse.Get().(*AlibabaIcbuRfqMyequityAPIResponse)
}

// ReleaseAlibabaIcbuRfqMyequityAPIResponse 将 AlibabaIcbuRfqMyequityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuRfqMyequityAPIResponse(v *AlibabaIcbuRfqMyequityAPIResponse) {
	v.Reset()
	poolAlibabaIcbuRfqMyequityAPIResponse.Put(v)
}
