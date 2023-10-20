package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletefaceAPIResponse 大麦换验平台-第三方对外开放-票面接口deleteFace API返回值
// alibaba.damai.mev.open.deleteface
//
// deleteFace
type AlibabaDamaiMevOpenDeletefaceAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletefaceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletefaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeletefaceAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeletefaceAPIResponseModel is 大麦换验平台-第三方对外开放-票面接口deleteFace 成功返回结果
type AlibabaDamaiMevOpenDeletefaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletefaceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletefaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeletefaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletefaceAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeletefaceAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletefaceAPIResponse
func GetAlibabaDamaiMevOpenDeletefaceAPIResponse() *AlibabaDamaiMevOpenDeletefaceAPIResponse {
	return poolAlibabaDamaiMevOpenDeletefaceAPIResponse.Get().(*AlibabaDamaiMevOpenDeletefaceAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeletefaceAPIResponse 将 AlibabaDamaiMevOpenDeletefaceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletefaceAPIResponse(v *AlibabaDamaiMevOpenDeletefaceAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletefaceAPIResponse.Put(v)
}
