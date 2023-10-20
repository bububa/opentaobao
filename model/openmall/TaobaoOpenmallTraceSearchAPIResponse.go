package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTraceSearchAPIResponse 获取Openmall订单物流流转信息 API返回值
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
type TaobaoOpenmallTraceSearchAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTraceSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTraceSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTraceSearchAPIResponseModel).Reset()
}

// TaobaoOpenmallTraceSearchAPIResponseModel is 获取Openmall订单物流流转信息 成功返回结果
type TaobaoOpenmallTraceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trace_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopLogisticsDetailTraceVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTraceSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTraceSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTraceSearchAPIResponse)
	},
}

// GetTaobaoOpenmallTraceSearchAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTraceSearchAPIResponse
func GetTaobaoOpenmallTraceSearchAPIResponse() *TaobaoOpenmallTraceSearchAPIResponse {
	return poolTaobaoOpenmallTraceSearchAPIResponse.Get().(*TaobaoOpenmallTraceSearchAPIResponse)
}

// ReleaseTaobaoOpenmallTraceSearchAPIResponse 将 TaobaoOpenmallTraceSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTraceSearchAPIResponse(v *TaobaoOpenmallTraceSearchAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTraceSearchAPIResponse.Put(v)
}
