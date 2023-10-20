package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse 装箱接单 API返回值
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
type AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel).Reset()
}

// AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel is 装箱接单 成功返回结果
type AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_load_container_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数
	Result *LoadReceiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse)
	},
}

// GetAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse 从 sync.Pool 获取 AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse
func GetAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse() *AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse {
	return poolAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse.Get().(*AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse)
}

// ReleaseAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse 将 AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse(v *AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse) {
	v.Reset()
	poolAlibabaTcwmsOutboundLoadContainerReceiveAPIResponse.Put(v)
}
