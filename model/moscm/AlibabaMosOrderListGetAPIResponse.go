package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderListGetAPIResponse 批量查询订单交易 API返回值
// alibaba.mos.order.list.get
//
// 批量查询交易信息
type AlibabaMosOrderListGetAPIResponse struct {
	model.CommonResponse
	AlibabaMosOrderListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOrderListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOrderListGetAPIResponseModel).Reset()
}

// AlibabaMosOrderListGetAPIResponseModel is 批量查询订单交易 成功返回结果
type AlibabaMosOrderListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_order_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaMosOrderListGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOrderListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOrderListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOrderListGetAPIResponse)
	},
}

// GetAlibabaMosOrderListGetAPIResponse 从 sync.Pool 获取 AlibabaMosOrderListGetAPIResponse
func GetAlibabaMosOrderListGetAPIResponse() *AlibabaMosOrderListGetAPIResponse {
	return poolAlibabaMosOrderListGetAPIResponse.Get().(*AlibabaMosOrderListGetAPIResponse)
}

// ReleaseAlibabaMosOrderListGetAPIResponse 将 AlibabaMosOrderListGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOrderListGetAPIResponse(v *AlibabaMosOrderListGetAPIResponse) {
	v.Reset()
	poolAlibabaMosOrderListGetAPIResponse.Put(v)
}
