package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteprojectAPIResponse 大麦换验平台-第三方对外开放-项目接口deleteProject API返回值
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
type AlibabaDamaiMevOpenDeleteprojectAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteprojectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteprojectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeleteprojectAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeleteprojectAPIResponseModel is 大麦换验平台-第三方对外开放-项目接口deleteProject 成功返回结果
type AlibabaDamaiMevOpenDeleteprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteprojectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeleteprojectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteprojectAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeleteprojectAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeleteprojectAPIResponse
func GetAlibabaDamaiMevOpenDeleteprojectAPIResponse() *AlibabaDamaiMevOpenDeleteprojectAPIResponse {
	return poolAlibabaDamaiMevOpenDeleteprojectAPIResponse.Get().(*AlibabaDamaiMevOpenDeleteprojectAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeleteprojectAPIResponse 将 AlibabaDamaiMevOpenDeleteprojectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeleteprojectAPIResponse(v *AlibabaDamaiMevOpenDeleteprojectAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeleteprojectAPIResponse.Put(v)
}
