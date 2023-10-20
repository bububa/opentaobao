package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderShowAPIResponse 闲鱼回收订单查询V1.1 API返回值
// alibaba.idle.recycle.order.show
//
// 查询回收订单
type AlibabaIdleRecycleOrderShowAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderShowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderShowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleOrderShowAPIResponseModel).Reset()
}

// AlibabaIdleRecycleOrderShowAPIResponseModel is 闲鱼回收订单查询V1.1 成功返回结果
type AlibabaIdleRecycleOrderShowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_show_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleRecycleOrderShowResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderShowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRecycleOrderShowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderShowAPIResponse)
	},
}

// GetAlibabaIdleRecycleOrderShowAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleOrderShowAPIResponse
func GetAlibabaIdleRecycleOrderShowAPIResponse() *AlibabaIdleRecycleOrderShowAPIResponse {
	return poolAlibabaIdleRecycleOrderShowAPIResponse.Get().(*AlibabaIdleRecycleOrderShowAPIResponse)
}

// ReleaseAlibabaIdleRecycleOrderShowAPIResponse 将 AlibabaIdleRecycleOrderShowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleOrderShowAPIResponse(v *AlibabaIdleRecycleOrderShowAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleOrderShowAPIResponse.Put(v)
}
