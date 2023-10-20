package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsTraceGetAPIResponse 用户根据交易号查询物流流转信息 API返回值
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
type TaobaoLogisticsTraceGetAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsTraceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsTraceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsTraceGetAPIResponseModel).Reset()
}

// TaobaoLogisticsTraceGetAPIResponseModel is 用户根据交易号查询物流流转信息 成功返回结果
type TaobaoLogisticsTraceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_trace_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result []TransitStepResult `json:"result,omitempty" xml:"result>transit_step_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsTraceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
}

var poolTaobaoLogisticsTraceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsTraceGetAPIResponse)
	},
}

// GetTaobaoLogisticsTraceGetAPIResponse 从 sync.Pool 获取 TaobaoLogisticsTraceGetAPIResponse
func GetTaobaoLogisticsTraceGetAPIResponse() *TaobaoLogisticsTraceGetAPIResponse {
	return poolTaobaoLogisticsTraceGetAPIResponse.Get().(*TaobaoLogisticsTraceGetAPIResponse)
}

// ReleaseTaobaoLogisticsTraceGetAPIResponse 将 TaobaoLogisticsTraceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsTraceGetAPIResponse(v *TaobaoLogisticsTraceGetAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsTraceGetAPIResponse.Put(v)
}
