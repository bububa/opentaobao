package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOutboundAPIResponse 零售商品发货 API返回值
// alibaba.mj.oc.outbound
//
// 用于接收发货的数据
type AlibabaMjOcOutboundAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcOutboundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcOutboundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcOutboundAPIResponseModel).Reset()
}

// AlibabaMjOcOutboundAPIResponseModel is 零售商品发货 成功返回结果
type AlibabaMjOcOutboundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_outbound_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcOutboundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaMjOcOutboundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcOutboundAPIResponse)
	},
}

// GetAlibabaMjOcOutboundAPIResponse 从 sync.Pool 获取 AlibabaMjOcOutboundAPIResponse
func GetAlibabaMjOcOutboundAPIResponse() *AlibabaMjOcOutboundAPIResponse {
	return poolAlibabaMjOcOutboundAPIResponse.Get().(*AlibabaMjOcOutboundAPIResponse)
}

// ReleaseAlibabaMjOcOutboundAPIResponse 将 AlibabaMjOcOutboundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcOutboundAPIResponse(v *AlibabaMjOcOutboundAPIResponse) {
	v.Reset()
	poolAlibabaMjOcOutboundAPIResponse.Put(v)
}
