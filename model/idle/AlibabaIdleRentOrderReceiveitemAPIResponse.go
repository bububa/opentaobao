package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderReceiveitemAPIResponse 确认签收 API返回值
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
type AlibabaIdleRentOrderReceiveitemAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentOrderReceiveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderReceiveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentOrderReceiveitemAPIResponseModel).Reset()
}

// AlibabaIdleRentOrderReceiveitemAPIResponseModel is 确认签收 成功返回结果
type AlibabaIdleRentOrderReceiveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_receiveitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderReceiveitemTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderReceiveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentOrderReceiveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderReceiveitemAPIResponse)
	},
}

// GetAlibabaIdleRentOrderReceiveitemAPIResponse 从 sync.Pool 获取 AlibabaIdleRentOrderReceiveitemAPIResponse
func GetAlibabaIdleRentOrderReceiveitemAPIResponse() *AlibabaIdleRentOrderReceiveitemAPIResponse {
	return poolAlibabaIdleRentOrderReceiveitemAPIResponse.Get().(*AlibabaIdleRentOrderReceiveitemAPIResponse)
}

// ReleaseAlibabaIdleRentOrderReceiveitemAPIResponse 将 AlibabaIdleRentOrderReceiveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentOrderReceiveitemAPIResponse(v *AlibabaIdleRentOrderReceiveitemAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentOrderReceiveitemAPIResponse.Put(v)
}
