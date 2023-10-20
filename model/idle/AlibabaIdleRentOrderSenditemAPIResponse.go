package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderSenditemAPIResponse 确认发货 API返回值
// alibaba.idle.rent.order.senditem
//
// 确认发货
type AlibabaIdleRentOrderSenditemAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentOrderSenditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderSenditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentOrderSenditemAPIResponseModel).Reset()
}

// AlibabaIdleRentOrderSenditemAPIResponseModel is 确认发货 成功返回结果
type AlibabaIdleRentOrderSenditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_senditem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderSenditemTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderSenditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentOrderSenditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderSenditemAPIResponse)
	},
}

// GetAlibabaIdleRentOrderSenditemAPIResponse 从 sync.Pool 获取 AlibabaIdleRentOrderSenditemAPIResponse
func GetAlibabaIdleRentOrderSenditemAPIResponse() *AlibabaIdleRentOrderSenditemAPIResponse {
	return poolAlibabaIdleRentOrderSenditemAPIResponse.Get().(*AlibabaIdleRentOrderSenditemAPIResponse)
}

// ReleaseAlibabaIdleRentOrderSenditemAPIResponse 将 AlibabaIdleRentOrderSenditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentOrderSenditemAPIResponse(v *AlibabaIdleRentOrderSenditemAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentOrderSenditemAPIResponse.Put(v)
}
