package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleCarOrderQueryAPIResponse 二手车寄卖查询订单接口 API返回值
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
type AlibabaIdleCarOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleCarOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleCarOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleCarOrderQueryAPIResponseModel).Reset()
}

// AlibabaIdleCarOrderQueryAPIResponseModel is 二手车寄卖查询订单接口 成功返回结果
type AlibabaIdleCarOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaIdleCarOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleCarOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleCarOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleCarOrderQueryAPIResponse)
	},
}

// GetAlibabaIdleCarOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleCarOrderQueryAPIResponse
func GetAlibabaIdleCarOrderQueryAPIResponse() *AlibabaIdleCarOrderQueryAPIResponse {
	return poolAlibabaIdleCarOrderQueryAPIResponse.Get().(*AlibabaIdleCarOrderQueryAPIResponse)
}

// ReleaseAlibabaIdleCarOrderQueryAPIResponse 将 AlibabaIdleCarOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleCarOrderQueryAPIResponse(v *AlibabaIdleCarOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleCarOrderQueryAPIResponse.Put(v)
}
