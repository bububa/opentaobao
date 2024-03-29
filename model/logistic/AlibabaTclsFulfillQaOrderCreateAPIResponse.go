package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsFulfillQaOrderCreateAPIResponse 创单接口 API返回值
// alibaba.tcls.fulfill.qa.order.create
//
// 根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
type AlibabaTclsFulfillQaOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTclsFulfillQaOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsFulfillQaOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsFulfillQaOrderCreateAPIResponseModel).Reset()
}

// AlibabaTclsFulfillQaOrderCreateAPIResponseModel is 创单接口 成功返回结果
type AlibabaTclsFulfillQaOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_fulfill_qa_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	RtErrorMsg string `json:"rt_error_msg,omitempty" xml:"rt_error_msg,omitempty"`
	// 错误码
	RtErrorCode int64 `json:"rt_error_code,omitempty" xml:"rt_error_code,omitempty"`
	// 是否成功
	RtSuccess bool `json:"rt_success,omitempty" xml:"rt_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsFulfillQaOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Data = ""
	m.RtErrorMsg = ""
	m.RtErrorCode = 0
	m.RtSuccess = false
}

var poolAlibabaTclsFulfillQaOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsFulfillQaOrderCreateAPIResponse)
	},
}

// GetAlibabaTclsFulfillQaOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaTclsFulfillQaOrderCreateAPIResponse
func GetAlibabaTclsFulfillQaOrderCreateAPIResponse() *AlibabaTclsFulfillQaOrderCreateAPIResponse {
	return poolAlibabaTclsFulfillQaOrderCreateAPIResponse.Get().(*AlibabaTclsFulfillQaOrderCreateAPIResponse)
}

// ReleaseAlibabaTclsFulfillQaOrderCreateAPIResponse 将 AlibabaTclsFulfillQaOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsFulfillQaOrderCreateAPIResponse(v *AlibabaTclsFulfillQaOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaTclsFulfillQaOrderCreateAPIResponse.Put(v)
}
