package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallSparepartsDetailsCreateAPIResponse 天猫蚁巢同步工单申请备件明细 API返回值
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
type AlibabaTmallSparepartsDetailsCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTmallSparepartsDetailsCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallSparepartsDetailsCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallSparepartsDetailsCreateAPIResponseModel).Reset()
}

// AlibabaTmallSparepartsDetailsCreateAPIResponseModel is 天猫蚁巢同步工单申请备件明细 成功返回结果
type AlibabaTmallSparepartsDetailsCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmall_spareparts_details_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	DisplayMessage string `json:"display_message,omitempty" xml:"display_message,omitempty"`
	// app名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误参数
	ErrorParams string `json:"error_params,omitempty" xml:"error_params,omitempty"`
	// 返回数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallSparepartsDetailsCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DisplayMessage = ""
	m.AppName = ""
	m.ErrorMessage = ""
	m.ErrorParams = ""
	m.Data = false
}

var poolAlibabaTmallSparepartsDetailsCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallSparepartsDetailsCreateAPIResponse)
	},
}

// GetAlibabaTmallSparepartsDetailsCreateAPIResponse 从 sync.Pool 获取 AlibabaTmallSparepartsDetailsCreateAPIResponse
func GetAlibabaTmallSparepartsDetailsCreateAPIResponse() *AlibabaTmallSparepartsDetailsCreateAPIResponse {
	return poolAlibabaTmallSparepartsDetailsCreateAPIResponse.Get().(*AlibabaTmallSparepartsDetailsCreateAPIResponse)
}

// ReleaseAlibabaTmallSparepartsDetailsCreateAPIResponse 将 AlibabaTmallSparepartsDetailsCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallSparepartsDetailsCreateAPIResponse(v *AlibabaTmallSparepartsDetailsCreateAPIResponse) {
	v.Reset()
	poolAlibabaTmallSparepartsDetailsCreateAPIResponse.Put(v)
}
