package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteFaceelementAPIResponse 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement API返回值
// alibaba.damai.mev.open.delete.faceelement
//
// deleteFaceElement
type AlibabaDamaiMevOpenDeleteFaceelementAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteFaceelementAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteFaceelementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeleteFaceelementAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeleteFaceelementAPIResponseModel is 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement 成功返回结果
type AlibabaDamaiMevOpenDeleteFaceelementAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_delete_faceelement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteFaceelementResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeleteFaceelementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeleteFaceelementAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeleteFaceelementAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeleteFaceelementAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeleteFaceelementAPIResponse
func GetAlibabaDamaiMevOpenDeleteFaceelementAPIResponse() *AlibabaDamaiMevOpenDeleteFaceelementAPIResponse {
	return poolAlibabaDamaiMevOpenDeleteFaceelementAPIResponse.Get().(*AlibabaDamaiMevOpenDeleteFaceelementAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeleteFaceelementAPIResponse 将 AlibabaDamaiMevOpenDeleteFaceelementAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeleteFaceelementAPIResponse(v *AlibabaDamaiMevOpenDeleteFaceelementAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeleteFaceelementAPIResponse.Put(v)
}
