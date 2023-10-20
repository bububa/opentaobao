package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrCepOrderQueryAPIResponse 订单详情查询接口 API返回值
// alibaba.health.nr.cep.order.query
//
// 订单详情查询接口
type AlibabaHealthNrCepOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrCepOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrCepOrderQueryAPIResponseModel).Reset()
}

// AlibabaHealthNrCepOrderQueryAPIResponseModel is 订单详情查询接口 成功返回结果
type AlibabaHealthNrCepOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_cep_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ResponseResult *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseResult = nil
}

var poolAlibabaHealthNrCepOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrCepOrderQueryAPIResponse)
	},
}

// GetAlibabaHealthNrCepOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaHealthNrCepOrderQueryAPIResponse
func GetAlibabaHealthNrCepOrderQueryAPIResponse() *AlibabaHealthNrCepOrderQueryAPIResponse {
	return poolAlibabaHealthNrCepOrderQueryAPIResponse.Get().(*AlibabaHealthNrCepOrderQueryAPIResponse)
}

// ReleaseAlibabaHealthNrCepOrderQueryAPIResponse 将 AlibabaHealthNrCepOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrCepOrderQueryAPIResponse(v *AlibabaHealthNrCepOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrCepOrderQueryAPIResponse.Put(v)
}
