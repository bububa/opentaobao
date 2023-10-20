package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPerfectPerformanceItemQueryAPIResponse 商品完美履约信息查询 API返回值
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
type AlibabaPerfectPerformanceItemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaPerfectPerformanceItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPerfectPerformanceItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPerfectPerformanceItemQueryAPIResponseModel).Reset()
}

// AlibabaPerfectPerformanceItemQueryAPIResponseModel is 商品完美履约信息查询 成功返回结果
type AlibabaPerfectPerformanceItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_perfect_performance_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *ItemPerfectPerformanceQueryResp `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPerfectPerformanceItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaPerfectPerformanceItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPerfectPerformanceItemQueryAPIResponse)
	},
}

// GetAlibabaPerfectPerformanceItemQueryAPIResponse 从 sync.Pool 获取 AlibabaPerfectPerformanceItemQueryAPIResponse
func GetAlibabaPerfectPerformanceItemQueryAPIResponse() *AlibabaPerfectPerformanceItemQueryAPIResponse {
	return poolAlibabaPerfectPerformanceItemQueryAPIResponse.Get().(*AlibabaPerfectPerformanceItemQueryAPIResponse)
}

// ReleaseAlibabaPerfectPerformanceItemQueryAPIResponse 将 AlibabaPerfectPerformanceItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPerfectPerformanceItemQueryAPIResponse(v *AlibabaPerfectPerformanceItemQueryAPIResponse) {
	v.Reset()
	poolAlibabaPerfectPerformanceItemQueryAPIResponse.Put(v)
}
