package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteperformAPIResponse 大麦换验平台-第三方对外开放-场次接口deletePerform API返回值
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
type AlibabaDamaiMevOpenDeleteperformAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteperformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteperformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeleteperformAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeleteperformAPIResponseModel is 大麦换验平台-第三方对外开放-场次接口deletePerform 成功返回结果
type AlibabaDamaiMevOpenDeleteperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteperform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteperformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteperformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeleteperformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteperformAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeleteperformAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeleteperformAPIResponse
func GetAlibabaDamaiMevOpenDeleteperformAPIResponse() *AlibabaDamaiMevOpenDeleteperformAPIResponse {
	return poolAlibabaDamaiMevOpenDeleteperformAPIResponse.Get().(*AlibabaDamaiMevOpenDeleteperformAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeleteperformAPIResponse 将 AlibabaDamaiMevOpenDeleteperformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeleteperformAPIResponse(v *AlibabaDamaiMevOpenDeleteperformAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeleteperformAPIResponse.Put(v)
}
