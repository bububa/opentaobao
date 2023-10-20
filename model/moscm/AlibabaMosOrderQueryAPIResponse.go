package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderQueryAPIResponse 批量查询订单信息 API返回值
// alibaba.mos.order.query
//
// 查询多笔交易信息
type AlibabaMosOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMosOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOrderQueryAPIResponseModel).Reset()
}

// AlibabaMosOrderQueryAPIResponseModel is 批量查询订单信息 成功返回结果
type AlibabaMosOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosOrderQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOrderQueryAPIResponse)
	},
}

// GetAlibabaMosOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaMosOrderQueryAPIResponse
func GetAlibabaMosOrderQueryAPIResponse() *AlibabaMosOrderQueryAPIResponse {
	return poolAlibabaMosOrderQueryAPIResponse.Get().(*AlibabaMosOrderQueryAPIResponse)
}

// ReleaseAlibabaMosOrderQueryAPIResponse 将 AlibabaMosOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOrderQueryAPIResponse(v *AlibabaMosOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaMosOrderQueryAPIResponse.Put(v)
}
