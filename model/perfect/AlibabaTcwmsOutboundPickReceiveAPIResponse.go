package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundPickReceiveAPIResponse 拣货接单 API返回值
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
type AlibabaTcwmsOutboundPickReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundPickReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundPickReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTcwmsOutboundPickReceiveAPIResponseModel).Reset()
}

// AlibabaTcwmsOutboundPickReceiveAPIResponseModel is 拣货接单 成功返回结果
type AlibabaTcwmsOutboundPickReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_pick_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *PickReceiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundPickReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTcwmsOutboundPickReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTcwmsOutboundPickReceiveAPIResponse)
	},
}

// GetAlibabaTcwmsOutboundPickReceiveAPIResponse 从 sync.Pool 获取 AlibabaTcwmsOutboundPickReceiveAPIResponse
func GetAlibabaTcwmsOutboundPickReceiveAPIResponse() *AlibabaTcwmsOutboundPickReceiveAPIResponse {
	return poolAlibabaTcwmsOutboundPickReceiveAPIResponse.Get().(*AlibabaTcwmsOutboundPickReceiveAPIResponse)
}

// ReleaseAlibabaTcwmsOutboundPickReceiveAPIResponse 将 AlibabaTcwmsOutboundPickReceiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTcwmsOutboundPickReceiveAPIResponse(v *AlibabaTcwmsOutboundPickReceiveAPIResponse) {
	v.Reset()
	poolAlibabaTcwmsOutboundPickReceiveAPIResponse.Put(v)
}
