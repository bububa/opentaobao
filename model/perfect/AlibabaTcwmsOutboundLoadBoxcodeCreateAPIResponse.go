package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse 创建箱号 API返回值
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
type AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel).Reset()
}

// AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel is 创建箱号 成功返回结果
type AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_load_boxcode_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BoxCodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse)
	},
}

// GetAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse 从 sync.Pool 获取 AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse
func GetAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse() *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse {
	return poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse.Get().(*AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse)
}

// ReleaseAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse 将 AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse(v *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse) {
	v.Reset()
	poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse.Put(v)
}
