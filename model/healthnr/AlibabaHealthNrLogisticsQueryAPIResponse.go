package healthnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsQueryAPIResponse 阿里健康新零售物流详情接口 API返回值
// alibaba.health.nr.logistics.query
//
// 对阿里健康o2o对接的商户提供查询物流单详情的能力
type AlibabaHealthNrLogisticsQueryAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrLogisticsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrLogisticsQueryAPIResponseModel).Reset()
}

// AlibabaHealthNrLogisticsQueryAPIResponseModel is 阿里健康新零售物流详情接口 成功返回结果
type AlibabaHealthNrLogisticsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_logistics_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResponseResult *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseResult = nil
}

var poolAlibabaHealthNrLogisticsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrLogisticsQueryAPIResponse)
	},
}

// GetAlibabaHealthNrLogisticsQueryAPIResponse 从 sync.Pool 获取 AlibabaHealthNrLogisticsQueryAPIResponse
func GetAlibabaHealthNrLogisticsQueryAPIResponse() *AlibabaHealthNrLogisticsQueryAPIResponse {
	return poolAlibabaHealthNrLogisticsQueryAPIResponse.Get().(*AlibabaHealthNrLogisticsQueryAPIResponse)
}

// ReleaseAlibabaHealthNrLogisticsQueryAPIResponse 将 AlibabaHealthNrLogisticsQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrLogisticsQueryAPIResponse(v *AlibabaHealthNrLogisticsQueryAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrLogisticsQueryAPIResponse.Put(v)
}
