package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectModifyAPIResponse 变更项目 API返回值
// alitrip.btrip.project.modify
//
// 变更项目
type AlitripBtripProjectModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripProjectModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripProjectModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripProjectModifyAPIResponseModel).Reset()
}

// AlitripBtripProjectModifyAPIResponseModel is 变更项目 成功返回结果
type AlitripBtripProjectModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_project_modify_response"`
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
func (m *AlitripBtripProjectModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Result = false
}

var poolAlitripBtripProjectModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripProjectModifyAPIResponse)
	},
}

// GetAlitripBtripProjectModifyAPIResponse 从 sync.Pool 获取 AlitripBtripProjectModifyAPIResponse
func GetAlitripBtripProjectModifyAPIResponse() *AlitripBtripProjectModifyAPIResponse {
	return poolAlitripBtripProjectModifyAPIResponse.Get().(*AlitripBtripProjectModifyAPIResponse)
}

// ReleaseAlitripBtripProjectModifyAPIResponse 将 AlitripBtripProjectModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripProjectModifyAPIResponse(v *AlitripBtripProjectModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripProjectModifyAPIResponse.Put(v)
}
