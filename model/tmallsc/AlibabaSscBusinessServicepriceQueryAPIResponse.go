package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscBusinessServicepriceQueryAPIResponse 苹果查询服务供给报价 API返回值
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
type AlibabaSscBusinessServicepriceQueryAPIResponse struct {
	model.CommonResponse
	AlibabaSscBusinessServicepriceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscBusinessServicepriceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscBusinessServicepriceQueryAPIResponseModel).Reset()
}

// AlibabaSscBusinessServicepriceQueryAPIResponseModel is 苹果查询服务供给报价 成功返回结果
type AlibabaSscBusinessServicepriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_business_serviceprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaSscBusinessServicepriceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscBusinessServicepriceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscBusinessServicepriceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscBusinessServicepriceQueryAPIResponse)
	},
}

// GetAlibabaSscBusinessServicepriceQueryAPIResponse 从 sync.Pool 获取 AlibabaSscBusinessServicepriceQueryAPIResponse
func GetAlibabaSscBusinessServicepriceQueryAPIResponse() *AlibabaSscBusinessServicepriceQueryAPIResponse {
	return poolAlibabaSscBusinessServicepriceQueryAPIResponse.Get().(*AlibabaSscBusinessServicepriceQueryAPIResponse)
}

// ReleaseAlibabaSscBusinessServicepriceQueryAPIResponse 将 AlibabaSscBusinessServicepriceQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscBusinessServicepriceQueryAPIResponse(v *AlibabaSscBusinessServicepriceQueryAPIResponse) {
	v.Reset()
	poolAlibabaSscBusinessServicepriceQueryAPIResponse.Put(v)
}
