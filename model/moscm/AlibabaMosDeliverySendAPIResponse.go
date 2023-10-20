package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosDeliverySendAPIResponse 发货 API返回值
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
type AlibabaMosDeliverySendAPIResponse struct {
	model.CommonResponse
	AlibabaMosDeliverySendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosDeliverySendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosDeliverySendAPIResponseModel).Reset()
}

// AlibabaMosDeliverySendAPIResponseModel is 发货 成功返回结果
type AlibabaMosDeliverySendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_delivery_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosDeliverySendResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosDeliverySendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosDeliverySendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosDeliverySendAPIResponse)
	},
}

// GetAlibabaMosDeliverySendAPIResponse 从 sync.Pool 获取 AlibabaMosDeliverySendAPIResponse
func GetAlibabaMosDeliverySendAPIResponse() *AlibabaMosDeliverySendAPIResponse {
	return poolAlibabaMosDeliverySendAPIResponse.Get().(*AlibabaMosDeliverySendAPIResponse)
}

// ReleaseAlibabaMosDeliverySendAPIResponse 将 AlibabaMosDeliverySendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosDeliverySendAPIResponse(v *AlibabaMosDeliverySendAPIResponse) {
	v.Reset()
	poolAlibabaMosDeliverySendAPIResponse.Put(v)
}
