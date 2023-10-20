package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushprojectAPIResponse 大麦换验平台-第三方对外开放-项目接口pushProject API返回值
// alibaba.damai.mev.open.pushproject
//
// pushProject
type AlibabaDamaiMevOpenPushprojectAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushprojectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushprojectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushprojectAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushprojectAPIResponseModel is 大麦换验平台-第三方对外开放-项目接口pushProject 成功返回结果
type AlibabaDamaiMevOpenPushprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushprojectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushprojectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushprojectAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushprojectAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushprojectAPIResponse
func GetAlibabaDamaiMevOpenPushprojectAPIResponse() *AlibabaDamaiMevOpenPushprojectAPIResponse {
	return poolAlibabaDamaiMevOpenPushprojectAPIResponse.Get().(*AlibabaDamaiMevOpenPushprojectAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushprojectAPIResponse 将 AlibabaDamaiMevOpenPushprojectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushprojectAPIResponse(v *AlibabaDamaiMevOpenPushprojectAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushprojectAPIResponse.Put(v)
}
