package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderQueryAPIResponse 闲鱼回收订单查询V1 API返回值
// alibaba.idle.recycle.order.query
//
// 查询回收订单
type AlibabaIdleRecycleOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleOrderQueryAPIResponseModel).Reset()
}

// AlibabaIdleRecycleOrderQueryAPIResponseModel is 闲鱼回收订单查询V1 成功返回结果
type AlibabaIdleRecycleOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleRecycleOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRecycleOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderQueryAPIResponse)
	},
}

// GetAlibabaIdleRecycleOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleOrderQueryAPIResponse
func GetAlibabaIdleRecycleOrderQueryAPIResponse() *AlibabaIdleRecycleOrderQueryAPIResponse {
	return poolAlibabaIdleRecycleOrderQueryAPIResponse.Get().(*AlibabaIdleRecycleOrderQueryAPIResponse)
}

// ReleaseAlibabaIdleRecycleOrderQueryAPIResponse 将 AlibabaIdleRecycleOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleOrderQueryAPIResponse(v *AlibabaIdleRecycleOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleOrderQueryAPIResponse.Put(v)
}
