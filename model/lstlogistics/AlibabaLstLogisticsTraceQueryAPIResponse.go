package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsTraceQueryAPIResponse 供应商-异云-查询运单物流追踪信息 API返回值
// alibaba.lst.logistics.trace.query
//
// 查询LP单物流追踪信息
type AlibabaLstLogisticsTraceQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstLogisticsTraceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsTraceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstLogisticsTraceQueryAPIResponseModel).Reset()
}

// AlibabaLstLogisticsTraceQueryAPIResponseModel is 供应商-异云-查询运单物流追踪信息 成功返回结果
type AlibabaLstLogisticsTraceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_trace_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstLogisticsTraceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsTraceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstLogisticsTraceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsTraceQueryAPIResponse)
	},
}

// GetAlibabaLstLogisticsTraceQueryAPIResponse 从 sync.Pool 获取 AlibabaLstLogisticsTraceQueryAPIResponse
func GetAlibabaLstLogisticsTraceQueryAPIResponse() *AlibabaLstLogisticsTraceQueryAPIResponse {
	return poolAlibabaLstLogisticsTraceQueryAPIResponse.Get().(*AlibabaLstLogisticsTraceQueryAPIResponse)
}

// ReleaseAlibabaLstLogisticsTraceQueryAPIResponse 将 AlibabaLstLogisticsTraceQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstLogisticsTraceQueryAPIResponse(v *AlibabaLstLogisticsTraceQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstLogisticsTraceQueryAPIResponse.Put(v)
}
