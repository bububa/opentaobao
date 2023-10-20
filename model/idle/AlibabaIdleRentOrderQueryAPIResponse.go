package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderQueryAPIResponse 查询订单 API返回值
// alibaba.idle.rent.order.query
//
// 查询订单信息
type AlibabaIdleRentOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentOrderQueryAPIResponseModel).Reset()
}

// AlibabaIdleRentOrderQueryAPIResponseModel is 查询订单 成功返回结果
type AlibabaIdleRentOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderQueryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderQueryAPIResponse)
	},
}

// GetAlibabaIdleRentOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleRentOrderQueryAPIResponse
func GetAlibabaIdleRentOrderQueryAPIResponse() *AlibabaIdleRentOrderQueryAPIResponse {
	return poolAlibabaIdleRentOrderQueryAPIResponse.Get().(*AlibabaIdleRentOrderQueryAPIResponse)
}

// ReleaseAlibabaIdleRentOrderQueryAPIResponse 将 AlibabaIdleRentOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentOrderQueryAPIResponse(v *AlibabaIdleRentOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentOrderQueryAPIResponse.Put(v)
}
