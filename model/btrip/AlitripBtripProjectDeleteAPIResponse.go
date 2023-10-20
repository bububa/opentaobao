package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectDeleteAPIResponse 删除项目 API返回值
// alitrip.btrip.project.delete
//
// 删除项目
type AlitripBtripProjectDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripProjectDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripProjectDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripProjectDeleteAPIResponseModel).Reset()
}

// AlitripBtripProjectDeleteAPIResponseModel is 删除项目 成功返回结果
type AlitripBtripProjectDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_project_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripProjectDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Result = false
}

var poolAlitripBtripProjectDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripProjectDeleteAPIResponse)
	},
}

// GetAlitripBtripProjectDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripProjectDeleteAPIResponse
func GetAlitripBtripProjectDeleteAPIResponse() *AlitripBtripProjectDeleteAPIResponse {
	return poolAlitripBtripProjectDeleteAPIResponse.Get().(*AlitripBtripProjectDeleteAPIResponse)
}

// ReleaseAlitripBtripProjectDeleteAPIResponse 将 AlitripBtripProjectDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripProjectDeleteAPIResponse(v *AlitripBtripProjectDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripProjectDeleteAPIResponse.Put(v)
}
